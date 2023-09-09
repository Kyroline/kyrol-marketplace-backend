package route

import (
	createProduct "kyrol-marketplace-backend/controllers/product-controllers/create"
	deleteProduct "kyrol-marketplace-backend/controllers/product-controllers/delete"
	getProduct "kyrol-marketplace-backend/controllers/product-controllers/get"
	updateProduct "kyrol-marketplace-backend/controllers/product-controllers/update"

	middleware "kyrol-marketplace-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func InitProductRoute(route *gin.Engine) {
	groupRoute := route.Group("/api/v1/product")
	groupRoute.Use(middleware.Auth())
	groupRoute.GET("/:id", middleware.Gate("product_read"), getProduct.ShowProduct)
	groupRoute.GET("/", middleware.Gate("product_read"), getProduct.GetProduct)
	groupRoute.POST("/", middleware.Gate("product_create"), createProduct.CreateProduct)
	groupRoute.PUT("/:id", middleware.Gate("product_update"), updateProduct.UpdateProduct)
	groupRoute.PATCH("/:id", middleware.Gate("product_update"), updateProduct.UpdateProduct)
	groupRoute.DELETE("/:id", middleware.Gate("product_delete"), deleteProduct.DeleteProduct)
}
