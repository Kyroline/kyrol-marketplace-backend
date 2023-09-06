package route

import (
	createCart "kyrol-marketplace-backend/controllers/cart-controllers/create"
	deleteCart "kyrol-marketplace-backend/controllers/cart-controllers/delete"
	getCart "kyrol-marketplace-backend/controllers/cart-controllers/get"
	updateCart "kyrol-marketplace-backend/controllers/cart-controllers/update"

	"github.com/gin-gonic/gin"
)

func InitCartRoute(route *gin.Engine) {
	group := route.Group("/api/v1/cart")
	group.GET("/", getCart.GetCart)
	group.POST("/", createCart.CreateCart)
	group.PATCH("/:id", updateCart.UpdateCart)
	group.DELETE("/:id", deleteCart.DeleteCart)
}
