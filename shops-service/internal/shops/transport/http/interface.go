package http

import (
	"context"
	"github.com/nahwinrajan/resilient-go/shops-service/domain"
)

type Usecases interface {
	GetShopProducts(ctx context.Context, shopID int64) ([]domain.Product, error)
}
