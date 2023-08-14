package createVariant

type Input struct {
	ProductID string `json:"product_id"`
	ID        string
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Stock     uint    `json:"stock"`
}
