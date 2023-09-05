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

	var cart Output
	if err := database.DB.Model(&model.Cart{}).Where("UserID = ?", uid).Take(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
		return
	}

	for i, item := range cart.Item {
		var product = model.Product{ID: item.ProductID}
		database.DB.Model(&model.Product{}).Find(&product)
		cart.Item[i].ProductName = product.Name
	}

	c.JSON(http.StatusOK, cart)
}
