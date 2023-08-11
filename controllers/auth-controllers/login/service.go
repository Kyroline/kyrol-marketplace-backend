package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	database "kyrol-marketplace-backend/database"
	model "kyrol-marketplace-backend/models"
	util "kyrol-marketplace-backend/utils"
)

func Login(c *gin.Context) {
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
		return
	}

	u := model.User{}

	if err := database.DB.Model(&model.User{}).Where("email = ?", input.Email).Take(&u).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid login credentials!",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid login credentials!",
		})
		return
	}

	token, err := util.GenerateToken(u.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"token": token,
	})
	return
}
