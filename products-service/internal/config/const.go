package config

import "fmt"

var configFilePaths = []string{
	"./products",
	"./files/etc/products-service/products",
	"/etc/products-service/products",
	"/files/etc/products-service/products",
	"../files/etc/products-service/products",
	"../../files/etc/products-service/products",
	"../../files/etc/products-service/products",
	"../../../files/etc/products-service/products",
}

var errFileNotFound = fmt.Errorf("Configuration file not found: %#v", configFilePaths)
