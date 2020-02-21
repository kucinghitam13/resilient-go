package http

import (
	"context"
	"github.com/nahwinrajan/resilient-go/products-service/domain"
)

type Usecases interface {
	GetProductByShopID(ctx context.Context, shopID int64) ([]domain.Product, error)
}
