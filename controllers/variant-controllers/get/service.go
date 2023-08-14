package getVariant

import (
	"net/http"

	"github.com/gin-gonic/gin"

	database "kyrol-marketplace-backend/database"
	model "kyrol-marketplace-backend/models"
)

func GetVariant(c *gin.Context) {
	var variant []Output
	if err := database.DB.Model(&model.ProductVariant{}).Find(&variant).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": variant,
	})
}
