package domain

type Product struct {
	ProductID   int64   `json:"product_id" db:"product_id"`
	ShopID      int64   `json:"shop_id" db:"shop_id"`
	Name        string  `json:"name" db:"name"`
	Description int64   `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	Stock       int64   `json:"stock" db:"stock"`
}
