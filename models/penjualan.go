package models

import "time"

// Penjualan represents a record in the Penjualan table.
type Penjualan struct {
	ID                int       `json:"id"`
	TransactionNumber string    `json:"transaction_number"`
	MarketingID       int       `json:"marketing_id"`
	Date              time.Time `json:"date"`
	CargoFee          float64   `json:"cargo_fee"`
	TotalBalance      float64   `json:"total_balance"`
	GrandTotal        float64   `json:"grand_total"`
}
