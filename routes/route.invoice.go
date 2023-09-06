package route

import (
	createInvoice "kyrol-marketplace-backend/controllers/invoice-controllers/create"
	getInvoice "kyrol-marketplace-backend/controllers/invoice-controllers/get"

	"github.com/gin-gonic/gin"
)

func InitInvoiceRoute(route *gin.Engine) {
	group := route.Group("/api/v1/invoice")
	group.GET("/", getInvoice.GetInvoice)
	group.POST("/", createInvoice.CreateInvoice)
}
