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

	id := "CART" + uid
	cart := model.Cart{UserID: uid, ID: id}

	if err := database.DB.Create(&cart); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
		return
	}

	c.Status(http.StatusNoContent)
}
