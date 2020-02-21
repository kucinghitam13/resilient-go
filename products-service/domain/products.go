package domain

// Products data model representation of Product entity
type Product struct {
	ProductID int64   `json:"product_id" db:"product_id"`
	ShopID    int64   `json:"shop_id" db:"product_id"`
	Name      string  `json:"name" db:"name"`
	Price     float64 `json:"price" db:"price"`
	Stock     int32   `json:"stock" db:"stock"`
}
