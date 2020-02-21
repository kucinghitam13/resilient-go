package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	"github.com/nahwinrajan/resilient-go/shops-service/internal/config"
	shops "github.com/nahwinrajan/resilient-go/shops-service/internal/shops"
	prdServiceRepo "github.com/nahwinrajan/resilient-go/shops-service/internal/repository/service_products"

)

var router *httprouter.Router

func main(){

	err := config.Init()
	if err != nil {
		log.Panicf("[main] error read config file :%+v\n", err)
	}

	serverConfig := config.GetServer()
	productsServiceConfig := config.GetServiceProducts()
	router = httprouter.New()
	clientHTTP := http.Client{
		Timeout: time.Duration(productsServiceConfig.HTTPTimeout) * time.Millisecond,
	}

	// initiate repositories (http)
	productServiceRepo := prdServiceRepo.New(productsServiceConfig, &clientHTTP)

	// initiate usecases
	shops.NewTransportHTTP(productServiceRepo, router)


	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", serverConfig.Port), router))
}