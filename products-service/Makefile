export PRODUCT_MAIN_FILE=./cmd/products-service
export PRODUCT_BINARY_PATH=./bin/
export PRODUCT_PKGS=$(shell go list ./... | grep -v vendor/)


build:
	@echo "${NOW} == build product service: ${PRODUCT_MAIN_FILE} ${PRODUCT_BINARY_PATH}"
	@CGO_ENABLED=0 go build -a -o ${PRODUCT_BINARY_PATH} ${PRODUCT_MAIN_FILE}