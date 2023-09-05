package getCategory

type Category struct {
	ID       uint       `json:"id"`
	Name     string     `json:"name"`
	Products []*Product `gorm:"many2many:category_product"`
}

type Product struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Categories  []*Category `gorm:"many2many:category_product" json:"-"`
}
