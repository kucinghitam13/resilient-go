export SHOP_MAIN_FILE=./cmd/shops-service
export SHOP_BINARY_PATH=./bin
export SHOP_PKGS=$(shell go list ./... | grep -v vendor/)

build:
	@echo "${NOW} == build shop service: ${SHOP_MAIN_FILE} ${SHOP_BINARY_PATH}"
	@CGO_ENABLED=0 go build -a -o ${SHOP_BINARY_PATH} ${SHOP_MAIN_FILE}