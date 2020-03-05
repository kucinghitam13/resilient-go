export PRODUCT_MAIN_FILE=./products-service/cmd/products-service
export PRODUCT_BINARY_PATH=./products-service/bin/
export PRODUCT_PKGS=$(shell go list ./products-service/... | grep -v vendor/)
export PRODUCT_FOLDER =./products-service/
export SHOP_MAIN_FILE=./shops-service/cmd/shops-service
export SHOP_BINARY_PATH=./shops-service/bin
export SHOP_PKGS=$(shell go list ./shops-service/... | grep -v vendor/)
export SHOPS_FOLDER =./shops-service/

test-products:
	@echo "testing products-service"
	@go test -cover -race ${PRODUCT_PKGS}

test-shops:
	@echo "testing shops-service"
	go test -cover -race ${SHOP_PKGS}

build-products:
	@echo "build product service: ${PRODUCT_MAIN_FILE} ${PRODUCT_BINARY_PATH}"
	@go build -a -o ${PRODUCT_BINARY_PATH} ${PRODUCT_MAIN_FILE}

build-shops:
	@echo "build product service: ${SHOP_MAIN_FILE} ${SHOP_BINARY_PATH}"
	@go build -a -o ${SHOP_BINARY_PATH} ${SHOP_MAIN_FILE}

docker-up:
	docker-compose up --build

docker-down:
	docker-compose down --rmi local