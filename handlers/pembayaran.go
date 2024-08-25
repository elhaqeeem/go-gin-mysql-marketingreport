package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/elhaqeeem/go-gin-mysql-marketingreport/models"
	"github.com/elhaqeeem/go-gin-mysql-marketingreport/utils"
	"github.com/gin-gonic/gin"
)

// handleCreditPayment handles payments in installments.
func handleCreditPayment(db *sql.DB, p *models.Pembayaran, jumlahAngsuran int) error {
	if p.Amount < float64(jumlahAngsuran) {
		return fmt.Errorf("jumlah angsuran tidak valid: harus lebih kecil dari total pembayaran")
	}

	// Calculate the installment amount.
	amountPerInstallment := p.Amount / float64(jumlahAngsuran)

	for i := 1; i <= jumlahAngsuran; i++ {
		// Set the installment payment date. Assuming each installment is paid monthly.
		tanggalPembayaran := p.PaymentDate.AddDate(0, i-1, 0)

		// Insert each installment into the PembayaranAngsuran table.
		if _, err := db.Exec(`
			INSERT INTO PembayaranAngsuran (PembayaranID, AngsuranKe, JumlahAngsuran, TanggalPembayaran)
			VALUES (?, ?, ?, ?)`,
			p.ID, i, amountPerInstallment, tanggalPembayaran,
		); err != nil {
			return err
		}
	}

	return nil
}

// CreatePembayaran handles payment creation, including installments.
func CreatePembayaran(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var p models.Pembayaran

		if err := c.ShouldBindJSON(&p); err != nil {
			utils.JSONResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		// Default value for PaymentDate if it is zero.
		if p.PaymentDate.IsZero() {
			p.PaymentDate = time.Now()
		}

		// Check if MarketingID exists.
		var exists bool
		err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM marketing WHERE ID = ?);", p.MarketingID).Scan(&exists)
		if err != nil {
			utils.JSONResponse(c, http.StatusInternalServerError, "database error")
			return
		}
		if !exists {
			utils.JSONResponse(c, http.StatusBadRequest, "marketingID does not exist")
			return
		}

		// Set status based on payment method.
		if p.PaymentMethod == "credit" {
			p.Status = "pending"
		} else {
			p.Status = "completed"
		}

		// Insert the main payment record.
		result, err := db.Exec(`
            INSERT INTO Pembayaran (MarketingID, Amount, PaymentDate, Status, PaymentMethod)
            VALUES (?, ?, ?, ?, ?)`,
			p.MarketingID, p.Amount, p.PaymentDate, p.Status, p.PaymentMethod,
		)
		if err != nil {
			utils.JSONResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		// Get the inserted payment ID.
		pembayaranID, err := result.LastInsertId()
		if err != nil {
			utils.JSONResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		p.ID = int(pembayaranID)

		// Handle the jumlah_angsuran query parameter.
		jumlahAngsuranStr := c.DefaultQuery("jumlah_angsuran", "1")
		jumlahAngsuran, err := strconv.Atoi(jumlahAngsuranStr)
		if err != nil {
			utils.JSONResponse(c, http.StatusBadRequest, "invalid jumlah_angsuran value")
			return
		}

		// Handle credit payments (installments).
		if p.PaymentMethod == "credit" {
			if err := handleCreditPayment(db, &p, jumlahAngsuran); err != nil {
				utils.JSONResponse(c, http.StatusBadRequest, err.Error())
				return
			}
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

		for rows.Next() {
			var p models.Pembayaran
			var paymentDateStr string

			if err := rows.Scan(&p.ID, &p.MarketingID, &p.Amount, &paymentDateStr, &p.Status, &p.PaymentMethod); err != nil {
				utils.JSONResponse(c, http.StatusInternalServerError, err.Error())
				return
			}

			// Adjust the time.Parse format string to match your database's date format (YYYY-MM-DD).
			p.PaymentDate, err = time.Parse("2006-01-02", paymentDateStr)
			if err != nil {
				utils.JSONResponse(c, http.StatusInternalServerError, "Invalid date format in database")
				return
			}

			payments = append(payments, p)
		}

		if err := rows.Err(); err != nil {
			utils.JSONResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, payments)
	}
}
