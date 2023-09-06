package getInvoice

import (
	"net/http"

	"github.com/gin-gonic/gin"

	database "kyrol-marketplace-backend/database"
	model "kyrol-marketplace-backend/models"
	util "kyrol-marketplace-backend/utils"
)

func GetInvoice(c *gin.Context) {
	id, _ := util.ExtractTokenID(c)

	var invoice []Output
	if err := database.DB.Model(&model.Invoice{}).Where("user_id  = ?", id).Preload("Items").Find(&invoice).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": invoice,
	})
}
