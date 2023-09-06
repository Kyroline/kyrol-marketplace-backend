package getInvoice

type Output struct {
	ID     string        `gorm:"type:varchar(30);primaryKey;not null"`
	UserID string        `gorm:"type:varchar(15)"`
	Items  []InvoiceItem `gorm:"foreignKey:InvoiceID"`
}

type InvoiceItem struct {
	InvoiceID string
	ID        string `gorm:"primaryKey"`
	ProductID string
	Name      string
	Price     float64
	Qty       uint
	Total     float64
}
