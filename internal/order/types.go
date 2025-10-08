package order

import "time"

type Order struct {
	ID          string    `json:"id"`
	CustomerID  string    `json:"customer_id"`
	TotalAmount string    `json:"total_amount"`
	Items       []string  `json:"items"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}
