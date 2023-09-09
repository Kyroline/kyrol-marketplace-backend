package middleware

import (
	"kyrol-marketplace-backend/database"
	model "kyrol-marketplace-backend/models"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go/v4"

	"github.com/gin-gonic/gin"
)

func Gate1(c *gin.Context, p string) bool {
	claim, _ := c.Get("user")

	var user model.User
	database.DB.Model(&model.User{}).Where("id = ?", claim.(jwt.MapClaims)["uid"]).Take(&user)

	var role []model.Role
	database.DB.Model(&user).Association("Roles").Find(&role)

	if database.DB.Model(&role).Where("name = ?", p).Association("Permissions").Count() < 1 {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "This action is unauthorized",
		})
		return false
	}
	return true
}

func Gate(p string) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		claim, _ := c.Get("user")

		var user model.User
		database.DB.Model(&model.User{}).Where("id = ?", claim.(jwt.MapClaims)["uid"]).Take(&user)

		var role []model.Role
		database.DB.Model(&user).Association("Roles").Find(&role)

		if database.DB.Model(&role).Where("name = ?", p).Association("Permissions").Count() < 1 {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "This action is unauthorized",
			})
			defer c.AbortWithStatus(http.StatusForbidden)
		} else {
			c.Next()
		}
	})
}
