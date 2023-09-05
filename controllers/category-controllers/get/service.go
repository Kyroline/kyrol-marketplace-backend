package getCategory

import (
	database "kyrol-marketplace-backend/database"
	model "kyrol-marketplace-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategory(c *gin.Context) {
	var categories []Category
	if err := database.DB.Model(&model.Category{}).Preload("Products").Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": categories,
	})
}

func ShowCategory(c *gin.Context) {
	id := c.Param("id")

	var category model.Category
	if err := database.DB.Model(&model.Category{}).Preload("Products").Where("id = ?", id).First(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": category.Products,
	})
}
