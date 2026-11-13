package product

/* // Order -
type Order struct {
	ID    int         `json:"id,omitempty"`
	Items []OrderItem `json:"items,omitempty"`
} */

/* // OrderItem -
type OrderItem struct {
	Product   Product `json:"product"`
	Quantity int    `json:"quantity"`
} */

/* // Product -
type Product struct {
	ID          int                `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Price       string             `json:"price"`
} */

// POST request
type ProductRequest struct {
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Price       float64 `json:"price"`
}

// GET response
type ProductResponse struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    Price       string `json:"price"`
}

// PATCH request - nil fields are omitted
type ProductUpdateRequest struct {
    Name        *string  `json:"name,omitempty"`
    Description *string  `json:"description,omitempty"`
    Price       *float64 `json:"price,omitempty"`
}