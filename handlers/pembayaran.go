package handlers

import (
	"database/sql"
	"net/http"

	"github.com/elhaqeeem/go-gin-mysql-marketingreport/models"
	"github.com/elhaqeeem/go-gin-mysql-marketingreport/utils"

	"github.com/gin-gonic/gin"
)

// CreatePembayaran handles payment creation.
func CreatePembayaran(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var p models.Pembayaran
		if err := c.BindJSON(&p); err != nil {
			utils.JSONResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		_, err := db.Exec(`
            INSERT INTO Pembayaran (marketing_id, amount, payment_date, status)
            VALUES (?, ?, ?, ?)`,
			p.MarketingID, p.Amount, p.PaymentDate, p.Status,
		)
		if err != nil {
			utils.JSONResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}

// GetPembayaran retrieves all payments.
func GetPembayaran(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT id, marketing_id, amount, payment_date, status FROM Pembayaran")
		if err != nil {
			utils.JSONResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		defer rows.Close()

		var payments []models.Pembayaran
		for rows.Next() {
			var p models.Pembayaran
			if err := rows.Scan(&p.ID, &p.MarketingID, &p.Amount, &p.PaymentDate, &p.Status); err != nil {
				utils.JSONResponse(c, http.StatusInternalServerError, err.Error())
				return
			}
			payments = append(payments, p)
		}

		c.JSON(http.StatusOK, payments)
	}
}
