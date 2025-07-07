package router

import (
	"log"
	"net/http"

	"github.com/cr1phy/fitly/internal/models"
	"github.com/gin-gonic/gin"
)

func status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func getProduct(c *gin.Context) {
	filter := c.Query("filter")

	result := []models.Product{}
	models.DB.Model(&models.Product{}).Where("name LIKE ?", filter+"%").Group("name").Find(&result)

	if len(result) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
	}
	c.JSON(http.StatusOK, result)
}

func getProductById(c *gin.Context) {
	id := c.Param("id")

	var p models.Product
	if err := models.DB.First(&p, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}
	c.JSON(http.StatusOK, p)
}

func addProduct(c *gin.Context) {
	var p models.Product
	if err := c.ShouldBindBodyWithJSON(&p); err != nil {
		log.Panicln("something went wrong with body:", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
		return
	}

	if err := models.DB.Create(&p); err != nil {
		log.Panicln("something went wrong with creating product:", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully created!"})
}

func getDishes(c *gin.Context) {
	filter := c.Query("filter")

	result := []models.Dish{}
	models.DB.Model(&models.Dish{}).Where("name LIKE ?", filter+"%").Group("name").Find(&result)

	if len(result) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
	}
	c.JSON(http.StatusOK, result)
}

func getDishById(c *gin.Context) {
	id := c.Param("id")

	var d models.Dish
	if err := models.DB.First(&d, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}
	c.JSON(http.StatusOK, d)
}

func addDish(c *gin.Context) {
	var d models.Dish
	if err := c.ShouldBindBodyWithJSON(&d); err != nil {
		log.Panicln("something went wrong with body:", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
		return
	}

	if err := models.DB.Create(&d); err != nil {
		log.Panicln("something went wrong with creating product:", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully created!"})
}

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", status)
	r.GET("/products", getProduct)
	r.GET("/product/:id", getProductById)
	r.POST("/products", addProduct)
	r.GET("/dishes", getDishes)
	r.GET("/dish/:id", getDishById)
	r.POST("/dishes", addDish)

	return r
}
