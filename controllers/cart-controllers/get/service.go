package getCart

import (
	"net/http"

	"github.com/gin-gonic/gin"

	database "kyrol-marketplace-backend/database"
	model "kyrol-marketplace-backend/models"
	util "kyrol-marketplace-backend/utils"
)

func GetCart(c *gin.Context) {
	uid, _ := util.ExtractTokenID(c)

	var cart []Output
	if err := database.DB.Model(&model.Cart{}).Preload("Product").Where("user_id = ?", uid).Take(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": cart,
	})
}

func ShowCart(c *gin.Context) {
	id := c.Param("id")

	var cart Output
	if err := database.DB.Model(&model.Cart{}).Preload("Product").Where("id = ?", id).Take(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": cart,
	})
	return
}
