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

	product := model.Product{
		ID:          input.ID,
		Name:        input.Name,
		Description: input.Description,
		Stock:       input.Stock,
		Price:       input.Price,
	}

	if err := database.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	var category []model.Category
	if err := database.DB.Model(&model.Category{}).Where("id IN ?", input.Categories).Find(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	database.DB.Model(&product).Association("Categories").Append(&category)
	c.JSON(http.StatusCreated, product)
}
