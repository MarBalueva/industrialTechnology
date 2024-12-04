package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	ID            int64   `gorm:"primaryKey;column:id" json:"id"`
	CategoryId    int64   `gorm:"column:categoryid" json:"categoryId"`
	Name          string  `gorm:"column:name" json:"name"`
	Description   string  `gorm:"column:description" json:"description"`
	Proteins      float64 `gorm:"column:proteins" json:"proteins"`
	Fats          float64 `gorm:"column:fats" json:"fats"`
	Carbohydrates float64 `gorm:"column:carbohydrates" json:"carbohydrates"`
	Calories      float64 `gorm:"column:calories" json:"calories"`
	UnWeight      float64 `gorm:"column:unweight" json:"unWeight"`
	Weight        int     `gorm:"column:weight" json:"weight"`
	Packaged      int     `gorm:"column:packaged" json:"packaged"`
	CountTypePack int     `gorm:"column:counttypepack" json:"countTypePack"`
	Cost          float64 `gorm:"column:cost" json:"cost"`
	OkeiId        int     `gorm:"column:okeiid" json:"okeiId"`
	InStore       bool    `gorm:"column:instore" json:"inStore"`
	PhotoLink     string  `gorm:"column:photolink" json:"photoLink"`
}

type ProductInBasket struct {
	ProductID int64 `gorm:"primaryKey;column:productid" json:"productId"`
	UserID    int64 `gorm:"primaryKey;column:userid" json:"userId"`
	Count     int   `gorm:"column:count" json:"count"`
}

// var products = []Product{
// 	{ID: "1", CategoryId: 1, Name: "Торт Аленка", Description: "Это уникальный по своему составу торт без единого грамма муки. Миндальный корж, молочный ганаш, шоколадный мусс, шоколадный бисквит без муки, ореховое безе", Cost: 1995.0},
// 	{ID: "2", CategoryId: 2, Name: "Пирожное муравейник", Description: "Рассыпчатое песочное печенье, нежная вареная сгущёнка со сливочным маслом, изюм и грецкий орех.", Cost: 89.0},
// 	{ID: "3", CategoryId: 3, Name: "Макарон Вишня", Description: "Невесомые пирожные из тончайшей миндальной муки и с божественным вкусом вишни.", Cost: 79.0},
// }

var jwtKey = []byte("my_secret_key")

type Credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	UserID   int64  `json:"userId"`
	jwt.StandardClaims
}

type Appuser struct {
	ID         int64     `gorm:"primaryKey;column:id" json:"id"`
	Login      string    `gorm:"column:login;size:32" json:"login"`
	Password   string    `gorm:"column:password;size:32" json:"password"`
	EmpId      int64     `gorm:"column:empId" json:"empId"`
	ClientId   int64     `gorm:"column:clientId" json:"clientId"`
	CreateDate time.Time `gorm:"column:createDate" json:"createDate"`
	IsActive   bool      `gorm:"column:isActive" json:"isActive"`
}

// var users = []Appuser{
// 	{Username: "user", Password: "password"},
// 	{Username: "user1", Password: "password1"},
// 	{Username: "user2", Password: "password2"},
// 	{Username: "user3", Password: "password3"},
// }

// var cart = []CartItem{}

var router = gin.Default()

var db *gorm.DB

func main() {
	//Инициализация БД
	initDB()

	// Логин
	router.POST("/login", login)

	// Рефреш токена
	router.POST("/refresh", refreshToken)

	protected := router.Group("/")
	protected.Use(authMiddleware())
	{
		// Получение всех продуктов
		protected.GET("/products", getProducts)

		// Получение всех продуктов с таймаутом
		protected.GET("productswithtimeout", getProductsWithTimeout)

		// Получение продукта по ID
		protected.GET("/products/:id", getProductByID)

		// Создание нового продукта
		protected.POST("/products", createProduct)

		// Обновление существующего продукта
		protected.PUT("/products/:id", updateProduct)

		// Удаление продукта
		protected.DELETE("/products/:id", deleteProduct)

		// Получение всех продуктов в корзине
		protected.GET("/cart", getCart)

		// Добавление продукта в корзину
		protected.POST("/cart", addToCart)

		// Удаление продукта из корзины
		protected.DELETE("/cart/:productId", deleteFromCart)
	}

	router.Run(":8080")
}

func handleError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{"error": message})
}

func initDB() {
	dsn := "host=localhost user=postgres password=1 dbname=bakery port=5000 sslmode=disable search_path=bakery"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	} else {
		log.Println("Database connected successfully")
	}

	db.AutoMigrate(&Product{})
	db.AutoMigrate(&Appuser{})
	db.AutoMigrate(&ProductInBasket{})
	log.Println("Database initialized")
}

func generateToken(username string, userID int64) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		Username: username,
		UserID:   userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func login(c *gin.Context) {
	var creds Credentials
	if err := c.BindJSON(&creds); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid request")
		return
	}

	var user Appuser
	if err := db.Where("login = ? AND password = ?", creds.Login, creds.Password).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			handleError(c, http.StatusUnauthorized, "Unauthorized")
		} else {
			handleError(c, http.StatusInternalServerError, "Database error")
		}
		return
	}

	token, err := generateToken(creds.Login, user.ID)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "Could not create token")
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorExpired != 0 {
					handleError(c, http.StatusUnauthorized, "Token expired")
					c.Abort()
					return
				}
			}
			handleError(c, http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		if !token.Valid {
			handleError(c, http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		c.Set("userId", claims.UserID)
		c.Next()
	}
}

func refreshToken(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				newToken, err := generateToken(claims.Username, claims.UserID)
				if err != nil {
					handleError(c, http.StatusInternalServerError, "Could not refresh token")
					return
				}
				c.JSON(http.StatusOK, gin.H{"token": newToken})
				return
			}
		}
		handleError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if !token.Valid {
		handleError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}
	handleError(c, http.StatusBadRequest, "Token is still valid")
}

func getProducts(c *gin.Context) {
	var products []Product
	var total int64

	//пагинация
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "3")

	//фильтры
	name := c.Query("name")
	category := c.Query("category")

	//сортировка
	sortField := c.DefaultQuery("sort", "id")
	sortOrder := c.DefaultQuery("order", "asc")

	//допустимые поля для сортировки
	validSortFields := map[string]bool{
		"id":       true,
		"name":     true,
		"category": true,
		"cost":     true,
	}

	//проверка параметров пагинации
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		handleError(c, http.StatusBadRequest, "Invalid page parameter")
		return
	}

	limitInt, err2 := strconv.Atoi(limit)
	if err2 != nil || limitInt < 1 {
		handleError(c, http.StatusBadRequest, "Invalid limit parameter")
		return
	}

	offset := (pageInt - 1) * limitInt

	//проверка корректности поля сортировки
	if !validSortFields[sortField] {
		handleError(c, http.StatusBadRequest, "Invalid sort field")
		return
	}

	//проверка корректности порядка сортировки
	if sortOrder != "asc" && sortOrder != "desc" {
		handleError(c, http.StatusBadRequest, "Invalid sort order")
		return
	}

	query := db.Model(&Product{}).Limit(limitInt).Offset(offset)

	//проверка фильтров
	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}
	if category != "" {
		query = query.Where("category ILIKE ?", "%"+category+"%")
	}

	query = query.Order(fmt.Sprintf("%s %s", sortField, sortOrder))

	if err := query.Count(&total).Error; err != nil {
		handleError(c, http.StatusInternalServerError, "Failed to count products")
		return
	}

	if err := query.Find(&products).Error; err != nil {
		handleError(c, http.StatusInternalServerError, "Failed to fetch products")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  products,
		"total": total,
		"page":  pageInt,
		"limit": limitInt,
		"sort": gin.H{
			"field": sortField,
			"order": sortOrder,
		},
		"filters": gin.H{
			"name":     name,
			"category": category,
		},
	})
}

func getProductsWithTimeout(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 1*time.Millisecond)
	defer cancel()

	var products []Product
	time.Sleep(10 * time.Millisecond)

	if err := db.WithContext(ctx).Find(&products).Error; err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			handleError(c, http.StatusRequestTimeout, "Request timed out")
		} else {
			handleError(c, http.StatusInternalServerError, "Failed to fetch products")
		}
		return
	}

	c.JSON(http.StatusOK, products)
}

func getProductByID(c *gin.Context) {
	id := c.Param("id")
	var product Product
	if err := db.First(&product, id).Error; err != nil {
		handleError(c, http.StatusNotFound, "Product not found")
		return
	}
	c.JSON(http.StatusOK, product)
}

func createProduct(c *gin.Context) {
	var newProduct Product
	if err := c.BindJSON(&newProduct); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid request")
		return
	}
	db.Create(&newProduct)
	c.JSON(http.StatusCreated, newProduct)
}

func updateProduct(c *gin.Context) {
	id := c.Param("id")
	var product Product

	//поиск продукта
	if err := db.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			handleError(c, http.StatusNotFound, "Product not found")
			return
		}
		handleError(c, http.StatusInternalServerError, "Failed to retrieve product")
		return
	}

	//обновление продукта
	if err := c.ShouldBindJSON(&product); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := db.Save(&product).Error; err != nil {
		handleError(c, http.StatusInternalServerError, "Failed to update product")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product updated successfully",
		"data":    product,
	})
}

func deleteProduct(c *gin.Context) {
	id := c.Param("id")

	if err := db.Delete(&Product{}, id).Error; err != nil {
		handleError(c, http.StatusNotFound, "Product not found")
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}

func getCart(c *gin.Context) {
	userID := c.GetInt64("userId")
	var cartItems []ProductInBasket
	if err := db.Where("userid = ?", userID).Find(&cartItems).Error; err != nil {
		handleError(c, http.StatusInternalServerError, "Failed to fetch cart")
		return
	}

	c.JSON(http.StatusOK, cartItems)
}

func addToCart(c *gin.Context) {
	var newItem ProductInBasket

	if err := c.BindJSON(&newItem); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid request")
		return
	}

	var existingItem ProductInBasket
	result := db.Where("productid = ? AND userid = ?", newItem.ProductID, newItem.UserID).First(&existingItem)

	if result.RowsAffected > 0 {
		existingItem.Count += newItem.Count
		if err := db.Save(&existingItem).Error; err != nil {
			handleError(c, http.StatusInternalServerError, "Failed to update cart item")
			return
		}
		c.JSON(http.StatusOK, existingItem)
		return
	}

	if err := db.Create(&newItem).Error; err != nil {
		handleError(c, http.StatusInternalServerError, "Failed to add item to cart")
		return
	}

	c.JSON(http.StatusCreated, newItem)
}

func deleteFromCart(c *gin.Context) {
	productID := c.Param("productId")
	userID := c.GetInt64("userId")

	var item ProductInBasket
	if err := db.Where("productid = ? AND userid = ?", productID, userID).First(&item).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			handleError(c, http.StatusNotFound, "Product not found in cart")
		} else {
			handleError(c, http.StatusInternalServerError, "Failed to fetch cart item")
		}
		return
	}

	if err := db.Delete(&item).Error; err != nil {
		handleError(c, http.StatusInternalServerError, "Failed to remove product from cart")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product removed from cart"})
}
