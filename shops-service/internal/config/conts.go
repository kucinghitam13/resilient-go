package config

import "fmt"

var configFilePaths = []string{
	"./shops",
	"./files/etc/shops-service/shops",
	"/etc/shops-service/shops",
	"/files/etc/shops-service/shops",
	"../files/etc/shops-service/shops",
	"../../files/etc/shops-service/shops",
	"../../../files/etc/shops-service/shops",
}

var errFileNotFound = fmt.Errorf("Configuration file not found: %#v", configFilePaths)
