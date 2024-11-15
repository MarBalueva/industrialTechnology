package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Product struct {
	ID          string  `json:"id"`
	CategoryId  int     `json:"categoryId"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Cost        float64 `json:"cost"`
}

type CartItem struct {
	ProductID string `json:"productId"`
	Quantity  int    `json:"quantity"`
}

var products = []Product{
	{ID: "1", CategoryId: 1, Name: "Торт Аленка", Description: "Это уникальный по своему составу торт без единого грамма муки. Миндальный корж, молочный ганаш, шоколадный мусс, шоколадный бисквит без муки, ореховое безе", Cost: 1995.0},
	{ID: "2", CategoryId: 2, Name: "Пирожное муравейник", Description: "Рассыпчатое песочное печенье, нежная вареная сгущёнка со сливочным маслом, изюм и грецкий орех.", Cost: 89.0},
	{ID: "3", CategoryId: 3, Name: "Макарон Вишня", Description: "Невесомые пирожные из тончайшей миндальной муки и с божественным вкусом вишни.", Cost: 79.0},
}

var jwtKey = []byte("my_secret_key")

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = []User{
	{Username: "user", Password: "password"},
	{Username: "user1", Password: "password1"},
	{Username: "user2", Password: "password2"},
	{Username: "user3", Password: "password3"},
}

var cart = []CartItem{}

var router = gin.Default()

func main() {
	// Логин
	router.POST("/login", login)

	// Рефреш токена
	router.POST("/refresh", refreshToken)

	protected := router.Group("/")
	protected.Use(authMiddleware())
	{
		// Получение всех продуктов
		protected.GET("/products", getProducts)

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

func generateToken(username string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Minute)
	claims := &Claims{
		Username: username,
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
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	var validUser *User
	for _, user := range users {
		if user.Username == creds.Username && user.Password == creds.Password {
			validUser = &user
			break
		}
	}

	if validUser == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	token, err := generateToken(creds.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create token"})
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
					c.JSON(http.StatusUnauthorized, gin.H{"message": "token expired"})
					c.Abort()
					return
				}
			}
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}

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
				newToken, err := generateToken(claims.Username)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"message": "could not refresh token"})
					return
				}
				c.JSON(http.StatusOK, gin.H{"token": newToken})
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "token is still valid"})
}

func getProducts(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

func getProductByID(c *gin.Context) {
	id := c.Param("id")

	for _, product := range products {
		if product.ID == id {
			c.JSON(http.StatusOK, product)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
}

func createProduct(c *gin.Context) {
	var newProduct Product

	if err := c.BindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	products = append(products, newProduct)
	c.JSON(http.StatusCreated, newProduct)
}

func updateProduct(c *gin.Context) {
	id := c.Param("id")
	var updatedProduct Product

	if err := c.BindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	for i, product := range products {
		if product.ID == id {
			products[i] = updatedProduct
			c.JSON(http.StatusOK, updatedProduct)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
}

func deleteProduct(c *gin.Context) {
	id := c.Param("id")

	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
}

func getCart(c *gin.Context) {
	c.JSON(http.StatusOK, cart)
}

func addToCart(c *gin.Context) {
	var newItem CartItem

	if err := c.BindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	for i, item := range cart {
		if item.ProductID == newItem.ProductID {
			cart[i].Quantity += newItem.Quantity
			c.JSON(http.StatusOK, cart[i])
			return
		}
	}

	cart = append(cart, newItem)
	c.JSON(http.StatusCreated, newItem)
}

func deleteFromCart(c *gin.Context) {
	productID := c.Param("productId")

	for i, item := range cart {
		if item.ProductID == productID {
			cart = append(cart[:i], cart[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "product removed from cart"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "product not found in cart"})
}
