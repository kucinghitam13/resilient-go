package http

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/nahwinrajan/resilient-go/shops-service/domain"
)

// Pong ping..pong
func (handler *HTTPHandler) Pong(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Fired Up, and Ready to GO!!")

	return
}

// GetShopProducts return products of respective shop
func (handler *HTTPHandler) GetShopProducts(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ctx := r.Context()

	// Setup header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	var response domain.ResponseHTTP
	shopID, err := strconv.ParseInt(params.ByName("shop_id"), 10, 64)
	if err != nil {
		response.Message = err.Error()
		response.Data = nil

		encoded, _ := json.Marshal(response)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(encoded)
		return
	}

	response.Data, err = handler.usecases.GetShopProducts(ctx, shopID)
	if err != nil {
		response.Message = err.Error()
		response.Data = nil

		encoded, _ := json.Marshal(response)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(encoded)
		return
	}

	encoded, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(encoded)
	return
}
