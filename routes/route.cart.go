package route

import (
	createCart "kyrol-marketplace-backend/controllers/cart-controllers/create"
	deleteCart "kyrol-marketplace-backend/controllers/cart-controllers/delete"
	getCart "kyrol-marketplace-backend/controllers/cart-controllers/get"
	updateCart "kyrol-marketplace-backend/controllers/cart-controllers/update"
	middleware "kyrol-marketplace-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func InitCartRoute(route *gin.Engine) {
	group := route.Group("/api/v1/cart")
	group.Use(middleware.Auth())
	group.GET("/", middleware.Gate("cart_read"), getCart.GetCart)
	group.POST("/", middleware.Gate("cart_create"), createCart.CreateCart)
	group.PATCH("/:id", middleware.Gate("cart_update"), updateCart.UpdateCart)
	group.PUT("/:id", middleware.Gate("cart_update"), updateCart.UpdateCart)
	group.DELETE("/:id", middleware.Gate("cart_delete"), deleteCart.DeleteCart)
}
