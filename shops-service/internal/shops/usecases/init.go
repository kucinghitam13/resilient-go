package usecases

// ShopInteractor entry point for usecases
type ShopInteractor struct {
	productRepo ProductServiceRepo
}

// New produce new interactor
func New(prdRepo ProductServiceRepo) *ShopInteractor {
	return &ShopInteractor{
		productRepo: prdRepo,
	}
}
