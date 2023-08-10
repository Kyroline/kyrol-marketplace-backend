package getProduct

import (
	database "kyrol-marketplace-backend/database"
	model "kyrol-marketplace-backend/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProduct(c *gin.Context) {
	var products []model.EntityProducts
	err := database.DB.Model(&model.EntityProducts{}).Preload("EntityVariants").Find(&products).Error
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"data": products,
		})
	}
}
