package getProduct

type Output struct {
	ID             string           `json:"id"`
	Name           string           `json:"name"`
	Description    string           `json:"description"`
	ProductVariant []ProductVariant `gorm:"foreignKey:ProductID;references:ID" json:"variant"`
}

type ProductVariant struct {
	ProductID string  `json:"-"`
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Stock     uint    `json:"stock"`
}
