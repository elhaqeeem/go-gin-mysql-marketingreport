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

// GetAngsuranDetail retrieves the details of a specific installment payment by ID.
func GetAngsuranDetail(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		angsuranIDStr := c.Param("id")
		angsuranID, err := strconv.Atoi(angsuranIDStr)
		if err != nil {
			utils.JSONResponse(c, http.StatusBadRequest, "invalid installment ID")
			return
		}

		type AngsuranDetail struct {
			ID                int       `json:"id"`
			PembayaranID      int       `json:"pembayaran_id"`
			AngsuranKe        int       `json:"angsuran_ke"`
			JumlahAngsuran    float64   `json:"jumlah_angsurans"`
			TanggalPembayaran time.Time `json:"tanggal_pembayaran"`
		}

		var detail AngsuranDetail
		var tanggalPembayaranStr string
		query := `
			SELECT ID, PembayaranID, AngsuranKe, JumlahAngsuran, TanggalPembayaran
			FROM PembayaranAngsuran
			WHERE ID = ?`
		if err := db.QueryRow(query, angsuranID).Scan(&detail.ID, &detail.PembayaranID, &detail.AngsuranKe, &detail.JumlahAngsuran, &tanggalPembayaranStr); err != nil {
			if err == sql.ErrNoRows {
				utils.JSONResponse(c, http.StatusNotFound, "installment not found")
			} else {
				utils.JSONResponse(c, http.StatusInternalServerError, err.Error())
			}
			return
		}

		// Correctly parse the date from the string.
		detail.TanggalPembayaran, err = time.Parse("2006-01-02", tanggalPembayaranStr)
		if err != nil {
			utils.JSONResponse(c, http.StatusInternalServerError, "invalid date format: "+tanggalPembayaranStr)
			return
		}

		// Respond with the installment details.
		c.JSON(http.StatusOK, detail)
	}
}

// GetAllAngsuran retrieves all installment payments associated with a specific payment ID.
func GetAllAngsuran(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		pembayaranIDStr := c.Param("pembayaran_id")
		pembayaranID, err := strconv.Atoi(pembayaranIDStr)
		if err != nil {
			utils.JSONResponse(c, http.StatusBadRequest, "invalid payment ID")
			return
		}

		type Angsuran struct {
			ID                int       `json:"id"`
			PembayaranID      int       `json:"pembayaran_id"`
			AngsuranKe        int       `json:"angsuran_ke"`
			JumlahAngsuran    float64   `json:"jumlah_angsurans"`
			TanggalPembayaran time.Time `json:"tanggal_pembayaran"`
		}

		var angsurans []Angsuran
		query := `
			SELECT ID, PembayaranID, AngsuranKe, JumlahAngsuran, TanggalPembayaran
			FROM PembayaranAngsuran
			WHERE PembayaranID = ?`

		rows, err := db.Query(query, pembayaranID)
		if err != nil {
			utils.JSONResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		defer rows.Close()

		for rows.Next() {
			var angsuran Angsuran
			var tanggalPembayaranStr string

			if err := rows.Scan(&angsuran.ID, &angsuran.PembayaranID, &angsuran.AngsuranKe, &angsuran.JumlahAngsuran, &tanggalPembayaranStr); err != nil {
				utils.JSONResponse(c, http.StatusInternalServerError, err.Error())
				return
			}

			// Parse the date from string
			angsuran.TanggalPembayaran, err = time.Parse("2006-01-02", tanggalPembayaranStr)
			if err != nil {
				utils.JSONResponse(c, http.StatusInternalServerError, "invalid date format: "+tanggalPembayaranStr)
				return
			}

			angsurans = append(angsurans, angsuran)
		}

		if err := rows.Err(); err != nil {
			utils.JSONResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		// Respond with the installment data.
		c.JSON(http.StatusOK, angsurans)
	}
}

// CheckInstallmentStatus checks the payment status of the first installment for a given payment ID.
func CheckInstallmentStatus(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		pembayaranIDStr := c.Param("pembayaran_id")

		// Convert the parameter to an integer.
		pembayaranID, err := strconv.Atoi(pembayaranIDStr)
		if err != nil {
			utils.JSONResponse(c, http.StatusBadRequest, "invalid payment ID")
			return
		}

		// Define the InstallmentStatus struct to include JumlahAngsuran.
		type InstallmentStatus struct {
			AngsuranKe        int     `json:"angsuran_ke"`
			JumlahAngsuran    float64 `json:"jumlah_angsuran"`
			TanggalPembayaran string  `json:"tanggal_pembayaran"` // Store as string first
			Status            string  `json:"status"`
		}

		var installment InstallmentStatus
		query := `
			SELECT AngsuranKe, JumlahAngsuran, TanggalPembayaran
			FROM PembayaranAngsuran
			WHERE PembayaranID = ? AND AngsuranKe = 1`

		row := db.QueryRow(query, pembayaranID)
		err = row.Scan(&installment.AngsuranKe, &installment.JumlahAngsuran, &installment.TanggalPembayaran)
		if err == sql.ErrNoRows {
			// If there are no rows, it means the installment has not been found.
			utils.JSONResponse(c, http.StatusNotFound, "installment not found")
			return
		} else if err != nil {
			utils.JSONResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		// Convert TanggalPembayaran string to time.Time
		tanggalPembayaran, err := time.Parse("2006-01-02", installment.TanggalPembayaran)
		if err != nil {
			utils.JSONResponse(c, http.StatusInternalServerError, "invalid date format")
			return
		}

		// Determine the status based on the payment date.
		if tanggalPembayaran.IsZero() {
			installment.Status = "belum dibayar" // Not paid
		} else {
			installment.Status = "telah dibayar" // Paid
		}

		c.JSON(http.StatusOK, installment) // Respond with installment status
	}
}
