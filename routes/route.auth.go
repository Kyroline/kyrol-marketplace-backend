package route

import (
	login "kyrol-marketplace-backend/controllers/auth-controllers/login"
	register "kyrol-marketplace-backend/controllers/auth-controllers/register"

	"github.com/gin-gonic/gin"
)

func InitAuthRoute(route *gin.Engine) {
	groupRoute := route.Group("/auth")
	groupRoute.POST("/login", login.Login)
	groupRoute.POST("/register", register.Register)
}
