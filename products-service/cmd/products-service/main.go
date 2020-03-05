package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/nahwinrajan/resilient-go/products-service/internal/config"
	products "github.com/nahwinrajan/resilient-go/products-service/internal/products"
	pgRepo "github.com/nahwinrajan/resilient-go/products-service/internal/repository/postgre"
)

var router *httprouter.Router

func main() {
	err := config.Init()
	if err != nil {
		log.Panicf("[main] error read config file :%+v\n", err)
	}

	serverConfig := config.GetServer()
	dbConfig := config.GetDatabase()
	router = httprouter.New()

	// initiate repo database
	pgDB, err := pgRepo.New(dbConfig)
	if err != nil {
		log.Panicf("[main] error connecto to database :%+v\n", err)
	}

	// initiate usecase
	products.NewTransportHTTP(pgDB, router)

	log.Printf("Products Service Fired Up and Ready to Go!")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", serverConfig.Port), router))
}
