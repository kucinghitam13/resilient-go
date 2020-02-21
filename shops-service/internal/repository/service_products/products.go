package service_products

import (
	"fmt"
	"context"
	"net/http"
	"encoding/json"

	"github.com/nahwinrajan/resilient-go/shops-service/domain"
)

// GetShopProducts return products owned by respective shop
func (sp *ServiceProducts) GetShopProducts(ctx context.Context, shopID int64) (products []domain.Product, err error) {
	products = make([]domain.Product, 0)

	req, err := http.NewRequestWithContext(ctx,
		http.MethodGet, 
		fmt.Sprintf("%s%s", 
			sp.conf.Host,
			fmt.Sprintf(sp.conf.URLGetProductsByShopID, shopID),
		),
		nil,
	)
	if err != nil {
		return products, err
	}
	req.Header.Set("Content-type", "application/json")
	
	resp, err := sp.do(req)
	if err != nil {
		return products, fmt.Errorf("service_products responded with error :%+v", err)
	}
	
	// resp.Body is an io.ReadCloser... NewDecoder expects an io.Reader
	var data domain.ResponseGetProductsByShop
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return products, fmt.Errorf("error read service_products response :%+v", err)
	}

	for _, val := range data.Products {
		products = append(products, val)
	}

	return products, nil
}