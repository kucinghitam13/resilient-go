export PRODUCT_MAIN_FILE=./products-service/cmd/products-service/main.go
export PRODUCT_BINARY_PATH=./products-service/bin/products-sevice
export PRODUCT_PKGS=$(shell go list ./products-service | grep -v vendor/)
export PRODUCT_FOLDER =./products-service/
export SHOP_MAIN_FILE=./shops-service/cmd/shops-service/main.go
export SHOP_BINARY_PATH=./shops-service/bin/shops-service
export SHOP_PKGS=$(shell go list ./shops-service | grep -v vendor/)
export SHOPS_FOLDER =./shops-service/

test-products:
	@echo "testing products-service"
	go test -cover -race ${PRODUCT_PKGS}

test-shops:
	@echo "testing shops-service"
	go test -cover -race ${SHOP_PKGS}

build-products:
	@echo "build product service: ${PRODUCT_MAIN_FILE} ${PRODUCT_BINARY_PATH}"
	go build -a -o ${PRODUCT_BINARY_PATH} ${PRODUCT_MAIN_FILE}

build-shops:
	go build -a -o ${SHOP_BINARY_PATH} ${SHOP_MAIN_FILE}

docker-up:
	build-products
	build-shops
	docker-compose up --build -d

docker-down:
	docker-compose down

