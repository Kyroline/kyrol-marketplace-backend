package route

import (
	createVariant "kyrol-marketplace-backend/controllers/variant-controllers/create"
	getVariant "kyrol-marketplace-backend/controllers/variant-controllers/get"

	"github.com/gin-gonic/gin"
)

func InitVariantRoute(route *gin.Engine) {
	groupRoute := route.Group("/api/v1/variant")
	groupRoute.GET("/", getVariant.GetVariant)
	groupRoute.POST("/", createVariant.CreateVariant)
}
