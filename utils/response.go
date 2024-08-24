package utils

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// JSONResponse is a utility function to send JSON responses.
func JSONResponse(c *gin.Context, statusCode int, message interface{}) {
	c.JSON(statusCode, gin.H{"error": message})
}

// GenerateTransactionNumber generates a new transaction number in the format TRX001, TRX002, etc.
func GenerateTransactionNumber(db *sql.DB) (string, error) {
	var lastTransactionNumber string
	query := "SELECT transaction_number FROM Penjualan ORDER BY ID DESC LIMIT 1"
	err := db.QueryRow(query).Scan(&lastTransactionNumber)
	if err != nil && err != sql.ErrNoRows {
		return "", fmt.Errorf("failed to retrieve last transaction number: %w", err)
	}

	if lastTransactionNumber == "" {
		return "TRX001", nil
	}

	// Extract the numeric part and increment it
	prefix := "TRX"
	numPart := lastTransactionNumber[len(prefix):]
	num, err := strconv.Atoi(numPart)
	if err != nil {
		return "", fmt.Errorf("failed to parse transaction number: %w", err)
	}

	num++
	return fmt.Sprintf("%s%03d", prefix, num), nil
}
