package product

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) GetProducts() ([]ProductResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/products", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	products := []ProductResponse{}
	err = json.Unmarshal(body, &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (c *Client) GetProduct(productID string) (*ProductResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/products/%s", c.HostURL, productID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	var product ProductResponse
	if err := json.Unmarshal(body, &product); err != nil {
        return nil, err
    }

	return &product, nil
}

func (c *Client) CreateProduct(product ProductRequest, authToken *string) (*ProductRequest, error) {
	rb, err := json.Marshal(product)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/products", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	newProduct := ProductRequest{}
	err = json.Unmarshal(body, &newProduct)
	if err != nil {
		return nil, err
	}

	return &newProduct, nil
}