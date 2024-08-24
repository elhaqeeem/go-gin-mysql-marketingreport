package models

// Pembayaran represents a record in the Pembayaran table.
type Pembayaran struct {
	ID          int     `json:"id"`
	MarketingID int     `json:"marketing_id"`
	Amount      float64 `json:"amount"`
	PaymentDate string  `json:"payment_date"`
	Status      string  `json:"status"`
}
