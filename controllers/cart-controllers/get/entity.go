package getCart

import "time"

type Output struct {
	ID        string
	Item      []Item `gorm:"foreignKey:CartID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Item struct {
	CartID      string `json:"-"`
	ProductID   string
	VariantID   string
	ProductName string
	VariantName string
	Qty         uint
}
