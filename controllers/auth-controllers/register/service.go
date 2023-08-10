package register

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	database "kyrol-marketplace-backend/database"
	model "kyrol-marketplace-backend/models"
)

func Register(c *gin.Context) {
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	u := model.EntityUsers{}
	u.FirstName = input.FirstName
	u.LastName = input.LastName
	u.Username = input.Username
	u.Email = input.Email
	u.Password = input.Password

	var count int64
	database.DB.Model(&model.EntityUsers{}).Count(&count)

	u.ID = "US" + strconv.FormatInt(7000000+count+1, 10)

	err := database.DB.Create(&u).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, "User created")
	return
}
