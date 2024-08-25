package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/elhaqeeem/go-gin-mysql-marketingreport/models"
	"github.com/elhaqeeem/go-gin-mysql-marketingreport/utils"
	"github.com/gin-gonic/gin"
)

// handleCreditPayment checks if the payment amount exceeds a specified limit for credit payments.
func handleCreditPayment(amount float64) error {
	if amount > 1000000 {
		return fmt.Errorf("amount %.2f exceeds the limit for credit payments", amount)
	}
	// Additional credit payment logic can go here...
	return nil
}

// CreatePembayaran handles payment creation.
func CreatePembayaran(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var p models.Pembayaran

		// Bind the JSON input to the Pembayaran struct.
		if err := c.ShouldBindJSON(&p); err != nil {
			utils.JSONResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		// Set default value for PaymentDate if it is zero.
		if p.PaymentDate.IsZero() {
			p.PaymentDate = time.Now() // Use the current time if no date is provided.
		}

		// Check if MarketingID exists
		var exists bool
		err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM marketing WHERE ID = ?)", p.MarketingID).Scan(&exists)
		if err != nil || !exists {
			utils.JSONResponse(c, http.StatusBadRequest, "MarketingID does not exist")
			return
		}

		// Perform credit payment checks if applicable.
		if p.PaymentMethod == "credit" {
			if err := handleCreditPayment(p.Amount); err != nil {
				utils.JSONResponse(c, http.StatusBadRequest, err.Error())
				return
			}
		}

		// Insert payment data into the database.
		if _, err := db.Exec(`
            INSERT INTO Pembayaran (MarketingID, Amount, PaymentDate, Status, PaymentMethod)
            VALUES (?, ?, ?, ?, ?)`,
			p.MarketingID, p.Amount, p.PaymentDate, p.Status, p.PaymentMethod,
		); err != nil {
			utils.JSONResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		// Respond with a success message.
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}

// GetPembayaran retrieves all payments from the database.
func GetPembayaran(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT id, MarketingID, Amount, PaymentDate, Status, PaymentMethod FROM Pembayaran")
		if err != nil {
			utils.JSONResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		defer rows.Close()

		var payments []models.Pembayaran

		// Scan each row into the Pembayaran struct.
		for rows.Next() {
			var p models.Pembayaran
			if err := rows.Scan(&p.ID, &p.MarketingID, &p.Amount, &p.PaymentDate, &p.Status, &p.PaymentMethod); err != nil {
				utils.JSONResponse(c, http.StatusInternalServerError, err.Error())
				return
			}
			payments = append(payments, p)
		}

		// Check for errors during row iteration.
		if err := rows.Err(); err != nil {
			utils.JSONResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		// Respond with the retrieved payments.
		c.JSON(http.StatusOK, payments)
	}
}
