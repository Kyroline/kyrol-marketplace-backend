package getCart

type Output struct {
	UserID    string
	ID        string
	ProductID string
	Qty       uint
	Product   Product `gorm:"foreignKey:ProductID"`
}

type Product struct {
	ID    string
	Name  string
	Stock uint
	Price float64
}
