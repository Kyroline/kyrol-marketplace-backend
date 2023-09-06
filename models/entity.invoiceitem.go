package model

type InvoiceItem struct {
	InvoiceID string
	ID        string `gorm:"primaryKey"`
	ProductID string
	Product   Product `gorm:"foreignKey:ProductID"`
	Name      string
	Price     float64
	Qty       uint
	Total     float64
}
