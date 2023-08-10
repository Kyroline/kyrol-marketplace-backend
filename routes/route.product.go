package route

import (
	createProduct "kyrol-marketplace-backend/controllers/product-controllers/create"
	deleteProduct "kyrol-marketplace-backend/controllers/product-controllers/delete"
	getProduct "kyrol-marketplace-backend/controllers/product-controllers/get"
	updateProduct "kyrol-marketplace-backend/controllers/product-controllers/update"

	"github.com/gin-gonic/gin"
)

func InitProductRoute(route *gin.Engine) {
	groupRoute := route.Group("/api/v1/product")
	groupRoute.GET("/", getProduct.GetProduct)
	groupRoute.POST("/", createProduct.CreateProduct)
	groupRoute.GET("/:id", getProduct.GetProduct)
	groupRoute.PUT("/:id", updateProduct.UpdateProduct)
	groupRoute.DELETE("/:id", deleteProduct.DeleteProduct)
}
