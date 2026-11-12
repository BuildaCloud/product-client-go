package product

// Order -
type Order struct {
	ID    int         `json:"id,omitempty"`
	Items []OrderItem `json:"items,omitempty"`
}

// OrderItem -
type OrderItem struct {
	Product   Product `json:"product"`
	Quantity int    `json:"quantity"`
}

// Product -
type Product struct {
	ID          int                `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Price       string             `json:"price"`
}