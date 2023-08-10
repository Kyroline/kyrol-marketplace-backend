package route

import (
	login "kyrol-marketplace-backend/controllers/auth-controllers/login"
	register "kyrol-marketplace-backend/controllers/auth-controllers/register"

	"github.com/gin-gonic/gin"
)

func InitAuthRoute(route *gin.Engine) {
	groupRoute := route.Group("/auth/user")
	groupRoute.GET("/login", login.Login)
	groupRoute.GET("/register", register.Register)
}
