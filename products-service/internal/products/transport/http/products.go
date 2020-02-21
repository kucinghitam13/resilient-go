package http

import (
	"io"
	"net/http"
	"encoding/json"

	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/nahwinrajan/resilient-go/products-service/domain"
)

func (handler *HttpHandler) Pong(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Fired Up, and Ready to GO!!")

	return
}

func (handler *HttpHandler) GetProductByShopID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ctx := r.Context()

	// Setup header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	var response domain.ResponseHttp
	shopID, err := strconv.ParseInt(params.ByName("shop_id"), 10, 64)
	if err != nil {
		response.Message = err.Error()
		response.Data = nil

		// Notes: for sake of simplicity, let's just focus on the resiliency implementation on this project
		encoded, _ := json.Marshal(response)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(encoded)
		return
	}

	// Notes: for sake of simplicity, let's just focus on the resiliency implementation on this project
	response.Data, err = handler.usecases.GetProductByShopID(ctx, shopID)
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