package createCart

import (
	"net/http"

	"github.com/gin-gonic/gin"

	database "kyrol-marketplace-backend/database"
	model "kyrol-marketplace-backend/models"
	util "kyrol-marketplace-backend/utils"
)

func CreateCart(c *gin.Context) {
	uid, err := util.ExtractTokenID(c)

	if err != nil {
		c.Status(http.StatusForbidden)
		return
	}

	var input InputCreateCart
	c.ShouldBindJSON(&input)

	cart := model.Cart{UserID: uid, ID: uid + input.ProductID, ProductID: input.ProductID, Qty: input.Qty}

	if err := database.DB.Create(&cart); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
		return
	}

	c.JSON(http.StatusCreated, cart)
}
