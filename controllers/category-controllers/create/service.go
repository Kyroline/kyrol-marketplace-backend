package createCategory

import (
	database "kyrol-marketplace-backend/database"
	model "kyrol-marketplace-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var input InputCreateCategory
	c.ShouldBindJSON(&input)

	category := model.Category{ID: input.ID, Name: input.Name}

	if err := database.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, input)
}
