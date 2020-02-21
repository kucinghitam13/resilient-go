package usecases

import (
	"context"
	"github.com/nahwinrajan/resilient-go/products-service/domain"
)

// PostgreRepo dependency injector for products repository
type PostgreRepo interface {
	GetProductsByShopID(ctx context.Context, shopID int64) ([]domain.Product, error)
}
