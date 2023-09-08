package createInvoice

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	database "kyrol-marketplace-backend/database"
	model "kyrol-marketplace-backend/models"
	util "kyrol-marketplace-backend/utils"
)

func CreateInvoice(c *gin.Context) {
	uid, err1 := util.ExtractTokenID(c)

	if err1 != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err1.Error(),
		})
		return
	}

	var input InputCreateInvoice
	c.ShouldBindJSON(&input)

	var invoice model.Invoice
	invoice.UserID = uid
	id := "INV-" + uid + "-" + time.Now().Format("20060102150405") //yyyyMMdHHmmss
	invoice.ID = id

	var items []model.InvoiceItem

	for p, item := range input.Items {
		var product model.Product
		if err := database.DB.Model(&model.Product{}).Where("id = ?", item.ProductID).Take(&product).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
		}

		var i model.InvoiceItem
		i.InvoiceID = id
		i.ID = id + strconv.Itoa(p+1)
		i.ProductID = item.ProductID
		i.Name = product.Name
		i.Price = product.Price
		i.Qty = item.Qty
		i.Total = product.Price * float64(item.Qty)

		items = append(items, i)
	}

	invoice.Items = items

	if err := database.DB.Create(&invoice).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, invoice)
}
