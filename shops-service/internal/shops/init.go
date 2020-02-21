package shops

import (
	"github.com/julienschmidt/httprouter"
	transHTTP "github.com/nahwinrajan/resilient-go/shops-service/internal/shops/transport/http"
	"github.com/nahwinrajan/resilient-go/shops-service/internal/shops/usecases"
)

// NewTransportHTTP setup http transport
func NewTransportHTTP(prdServiceRepo usecases.ProductServiceRepo, router *httprouter.Router) {
	ucs := usecases.New(prdServiceRepo)
	handler := transHTTP.New(router, ucs)
	handler.SetEndPoint()
}
