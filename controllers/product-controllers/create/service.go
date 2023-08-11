package createProduct

import (
	database "kyrol-marketplace-backend/database"
	model "kyrol-marketplace-backend/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var input InputCreateProduct
	c.ShouldBindJSON(&input)

	product := model.Product{ID: input.ID, Name: input.Name, Description: input.Description}

	result := database.DB.Create(&product)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Data is not created",
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{})
	}
}
