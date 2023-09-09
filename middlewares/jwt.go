package middleware

import (
	util "kyrol-marketplace-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {

	return gin.HandlerFunc(func(c *gin.Context) {
		if err := util.VerifyTokenHeader(c); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthenticated",
			})
			defer c.AbortWithStatus(http.StatusUnauthorized)
		} else {
			token, _ := util.ParseToken(util.ExtractTokenHeader(c))
			c.Set("user", token.Claims)
			c.Next()
		}
	})
}
