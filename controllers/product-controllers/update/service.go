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
	product.Stock = input.Stock
	product.Price = input.Price

	if err := database.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
		return
	}

	var category []model.Category
	if err := database.DB.Model(&model.Category{}).Where("id IN ?", input.Categories).Find(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	database.DB.Model(&product).Association("Categories").Replace(&category)

	c.Status(http.StatusNoContent)
}
