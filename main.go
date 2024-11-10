package main

import (
	"net/http"

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

var cart = []CartItem{}

var router = gin.Default()

func main() {
	// Получение всех продуктов
	router.GET("/products", getProducts)

	// Получение продукта по ID
	router.GET("/products/:id", getProductByID)

	// Создание нового продукта
	router.POST("/products", createProduct)

	// Обновление существующего продукта
	router.PUT("/products/:id", updateProduct)

	// Удаление продукта
	router.DELETE("/products/:id", deleteProduct)

	// Получение всех продуктов в корзине
	router.GET("/cart", getCart)

	// Добавление продукта в корзину
	router.POST("/cart", addToCart)

	// Удаление продукта из корзины
	router.DELETE("/cart/:productId", deleteFromCart)

	router.Run(":8080")
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
