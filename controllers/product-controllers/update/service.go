package updateProduct

import (
	database "kyrol-marketplace-backend/database"
	model "kyrol-marketplace-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	var input Input
	c.ShouldBindJSON(&input)

	var product model.Product
	if err := database.DB.Model(&model.Product{}).Where("ID = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
		return
	}

	product.Name = input.Name
	product.Description = input.Description

	if err := database.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
		return
	}

	c.Status(http.StatusNoContent)
}
