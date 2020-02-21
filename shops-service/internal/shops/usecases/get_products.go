package usecases

import (
	"context"

	"github.com/nahwinrajan/resilient-go/shops-service/domain"
)

// GetShopProducts get products of respective shop
func (interactor *ShopInteractor) GetShopProducts(ctx context.Context, shopID int64) ([]domain.Product, error) {
	products, err := interactor.productRepo.GetShopProducts(ctx, shopID)

	return products, err
}
