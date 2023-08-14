package createVariant

import (
	"net/http"

	"github.com/gin-gonic/gin"

	database "kyrol-marketplace-backend/database"
	model "kyrol-marketplace-backend/models"
)

func CreateVariant(c *gin.Context) {
	var input Input
	c.ShouldBindJSON(&input)

	var id string
	if input.ProductID != "" {
		id = input.ProductID
	} else if c.Param("id") != "" {
		id = c.Param("id")
	}

	variant := model.ProductVariant{ProductID: id, ID: input.ID, Name: input.Name, Price: input.Price, Stock: input.Stock}
	if err := database.DB.Save(&variant).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}