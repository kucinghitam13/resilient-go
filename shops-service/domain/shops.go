package domain

type Shops struct {
	ShopID int64  `json:"shop_id" db:"shop_id"`
	Name   string `json:"name" db:"name"`
	City   string `json:"city" db:"city"`
}
