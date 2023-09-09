package getProduct

import (
	database "kyrol-marketplace-backend/database"
	model "kyrol-marketplace-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProduct(c *gin.Context) {
	var products []Product

	err := database.DB.Model(&model.Product{}).Preload("Categories").Find(&products).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}

func ShowProduct(c *gin.Context) {
	id := c.Param("id")

	var product Product

	if err := database.DB.Model(&model.Product{}).Preload("Categories").Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}
