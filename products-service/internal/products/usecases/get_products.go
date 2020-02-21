package usecases

import (
	"context"
	"github.com/nahwinrajan/resilient-go/products-service/domain"
)

func (interactor *ProductInteractor) GetProductByShopID(ctx context.Context, shopID int64) ([]domain.Product, error) {
	products, err := interactor.pgRepo.GetProductsByShopID(ctx, shopID)

	return products, err
}
