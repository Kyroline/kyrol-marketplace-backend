package updateCategory

import (
	database "kyrol-marketplace-backend/database"
	model "kyrol-marketplace-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateCategory(c *gin.Context) {
	id := c.Param("id")

	var input InputUpdateCategory
	c.ShouldBindJSON(&input)

	var category model.Category
	if err := database.DB.Model(&model.Category{}).Where("id = ?", id).First(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	category.Name = input.Name
	if err := database.DB.Save(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
}
