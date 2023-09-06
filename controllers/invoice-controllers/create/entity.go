package createInvoice

type InputCreateInvoice struct {
	UserID string
	ID     string
	Items  []InvoiceItem
}

type InvoiceItem struct {
	ProductID string
	Qty       uint
}
