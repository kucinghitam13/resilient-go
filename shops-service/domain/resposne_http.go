package domain

// ResponseHTTP our service HTTP response structure
type ResponseHTTP struct {
	Message string
	Data interface{}
}

// ResponseGetProductsByShop response structure that will be returned
// from Products Server for get products by shop API
type ResponseGetProductsByShop struct {
	Message string
	Products []Product
}