package createItem

import (
	database "kyrol-marketplace-backend/database"
	model "kyrol-marketplace-backend/models"
	util "kyrol-marketplace-backend/utils"

	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateItem(c *gin.Context) {
	uid, _ := util.ExtractTokenID(c)

	var input Input
	c.ShouldBindJSON(&input)

	var cart model.Cart
	database.DB.Model(&model.Cart{}).Where("user_id = ?", uid).Take(&cart)

	var item model.CartItem
	item.CartID = cart.ID
	item.ID = cart.ID + input.ProductID
	item.ProductID = input.ProductID
	item.Qty = input.Qty

	if err := database.DB.Save(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}
}
