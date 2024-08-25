package models

import "time"

// Pembayaran represents a payment record.
type Pembayaran struct {
	ID            int       `json:"id"`
	MarketingID   int       `json:"marketing_id"`
	Amount        float64   `json:"amount"` // Change to float64
	PaymentDate   time.Time `json:"payment_date"`
	Status        string    `json:"status"`
	PaymentMethod string    `json:"PaymentMethod"`
}
