package deleteProduct

import (
	database "kyrol-marketplace-backend/database"
	model "kyrol-marketplace-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	var product model.Product
	if err := database.DB.Model(&model.Product{}).Where("ID = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error,
		})
		return
	}

	if err := database.DB.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
		return
	}

	c.Status(http.StatusNoContent)
}
