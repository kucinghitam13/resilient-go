package usecases

import (
	"context"
	"github.com/nahwinrajan/resilient-go/shops-service/domain"
)

// ProductServiceRepo interface for product-service repository layer
type ProductServiceRepo interface {
	GetShopProducts(ctx context.Context, shopID int64) ([]domain.Product, error)
}