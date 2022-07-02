package model

type Message struct {
	ID         string  `json:"id"`
	Total      float32 `json:"total"`
	Currency   string  `json:"currency"`
	CustomerID string  `json:"customerId"`
}
