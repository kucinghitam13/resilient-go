package http

import(
	"github.com/julienschmidt/httprouter"	
)


type HttpHandler struct {
	router *httprouter.Router
	usecases Usecases
}


func New(router *httprouter.Router, usecases Usecases) *HttpHandler {
	return &HttpHandler{
		router: router,
		usecases: usecases,
	}
}


func (handler *HttpHandler) SetEndPoint() {
	// NOTE: remember we are talking about resiliency here, not really proper web service project
	handler.router.GET("/ping", handler.Pong)
	// NOTE: yeah, didn't really gave careful thought about the sample entity, hence the api not that restful
	handler.router.GET("/api/v1/products/shop/:shop_id", handler.GetProductByShopID)
}