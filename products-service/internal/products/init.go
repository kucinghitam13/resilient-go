package products

import (
	"github.com/julienschmidt/httprouter"
	"github.com/nahwinrajan/resilient-go/products-service/internal/products/usecases"
	transHTTP "github.com/nahwinrajan/resilient-go/products-service/internal/products/transport/http"
)


func NewTransportHTTP(pgRepo usecases.PostgreRepo, router *httprouter.Router) {
	ucs := usecases.New(pgRepo)

	handler := transHTTP.New(router, ucs)
	handler.SetEndPoint()
}