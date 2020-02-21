package postgre

import (
	"context"

	"github.com/nahwinrajan/resilient-go/products-service/domain"
)

// GetProductsByShopID return products of respective shop from database
func (dbObj *DB) GetProductsByShopID(ctx context.Context, productID int64) ([]domain.Product, error) {
	dummyResult := make([]domain.Product, 0)
	query := `SELECT
		product_id,
		shop_id,
		name,
		description,
		price,
		stock
	FROM products
	WHERE shop_id = $1`

	rows, err := dbObj.queryxContext(ctx, query, productID)
	if err != nil {
		return dummyResult, err
	}
	defer rows.Close()

	products := make([]domain.Product, 0)
	for rows.Next() {
		var prod domain.Product
		err = rows.StructScan(&prod)
		if err != nil {
			return dummyResult, err
		}

		products = append(products, prod)
	}

	return products, nil
}
