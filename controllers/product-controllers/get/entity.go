package getProduct

type Product struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Categories  []*Category `gorm:"many2many:category_product"`
}

type Category struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Products []*Product `gorm:"many2many:category_product" json:"-"`
}
