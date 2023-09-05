package deleteCategory

import (
	database "kyrol-marketplace-backend/database"
	model "kyrol-marketplace-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	var category model.Category
	if err := database.DB.Model(&model.Category{}).Where("id = ?", id).First(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := database.DB.Delete(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
}
