package http

import (
	"github.com/julienschmidt/httprouter"
)

// HTTPHandler struct entry point for http transport
type HTTPHandler struct {
	router   *httprouter.Router
	usecases Usecases
}

// New return new http router
func New(router *httprouter.Router, usecases Usecases) *HTTPHandler {
	return &HTTPHandler{
		router:   router,
		usecases: usecases,
	}
}

// SetEndPoint register endpoints to http router
func (handler *HTTPHandler) SetEndPoint() {
	handler.router.GET("/ping", handler.Pong)
	handler.router.GET("/api/v1/shops/:shop_id/products", handler.GetShopProducts)
}
