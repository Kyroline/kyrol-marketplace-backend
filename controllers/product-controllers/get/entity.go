package getProduct

import "time"

type Product struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Categories  []*Category `gorm:"many2many:category_product"`
	Price       float64
	Stock       uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Category struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Products []*Product `gorm:"many2many:category_product" json:"-"`
}
