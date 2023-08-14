package getVariant

type Output struct {
	ProductID string  `json:"product_id"`
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Stock     uint    `json:"stock"`
}
