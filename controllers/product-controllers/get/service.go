package getProduct

import (
	database "kyrol-marketplace-backend/database"
	model "kyrol-marketplace-backend/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProduct(c *gin.Context) {
	var products []Output
	err := database.DB.Model(&model.Product{}).Preload("ProductVariant").Find(&products).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}
