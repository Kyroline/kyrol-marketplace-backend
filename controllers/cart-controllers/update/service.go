package updateCart

import (
	"net/http"

	"github.com/gin-gonic/gin"

	database "kyrol-marketplace-backend/database"
	model "kyrol-marketplace-backend/models"
)

func UpdateCart(c *gin.Context) {
	id := c.Param("id")

	var input InputUpdateCart
	c.ShouldBindJSON(&input)

	var cart model.Cart
	if err := database.DB.Model(&model.Cart{}).Where("id = ?", id).Take(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	cart.Qty = input.Qty
	if err := database.DB.Save(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Status(http.StatusAccepted)
}
