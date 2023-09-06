package createProduct

type InputCreateProduct struct {
	ID          string
	Name        string
	Description string
	Stock       uint
	Price       float64
	Categories  []uint
}
