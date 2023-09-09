package route

import (
	createInvoice "kyrol-marketplace-backend/controllers/invoice-controllers/create"
	getInvoice "kyrol-marketplace-backend/controllers/invoice-controllers/get"
	middleware "kyrol-marketplace-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func InitInvoiceRoute(route *gin.Engine) {
	group := route.Group("/api/v1/invoice")
	group.Use(middleware.Auth())
	group.GET("/", middleware.Gate("invoice_read"), getInvoice.GetInvoice)
	group.POST("/", middleware.Gate("invoice_create"), createInvoice.CreateInvoice)
}
