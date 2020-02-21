package usecases

// ProductInteractor structure for holding necessa
type ProductInteractor struct {
	pgRepo PostgreRepo
}

// New to generate object for accessing Products usecases
func New(postgreRepo PostgreRepo) *ProductInteractor {
	return &ProductInteractor{
		pgRepo: postgreRepo,
	}
}
