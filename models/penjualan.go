package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

// Penjualan represents a record in the Penjualan table.
type Penjualan struct {
	ID                int        `json:"id"`
	TransactionNumber string     `json:"TransactionNumber"`
	MarketingID       int        `json:"MarketingID"`
	Date              CustomTime `json:"Date"`
	CargoFee          float64    `json:"CargoFee"`
	TotalBalance      float64    `json:"TotalBalance"`
	GrandTotal        float64    `json:"GrandTotal"`
}

// Custom type for handling time parsing and JSON
type CustomTime struct {
	time.Time
}

// Implement UnmarshalJSON to parse the date string
func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	var dateStr string
	if err := json.Unmarshal(b, &dateStr); err != nil {
		return err
	}

	// Parse the date string
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return err
	}

	ct.Time = t
	return nil
}

// Implement Value method for driver.Valuer interface
func (ct CustomTime) Value() (driver.Value, error) {
	return ct.Time, nil // Convert to time.Time
}

// Implement Scan method for sql.Scanner interface
func (ct *CustomTime) Scan(value interface{}) error {
	if value == nil {
		ct.Time = time.Time{} // Set to zero time
		return nil
	}

	// Check for []byte type and convert to time.Time
	switch v := value.(type) {
	case []byte:
		// Try to parse the byte slice into time
		t, err := time.Parse("2006-01-02", string(v))
		if err != nil {
			return err // Return error if parsing fails
		}
		ct.Time = t
		return nil
	case time.Time:
		ct.Time = v // Directly assign if it's already time.Time
		return nil
	default:
		return fmt.Errorf("cannot scan type %T into CustomTime", value)
	}
}
