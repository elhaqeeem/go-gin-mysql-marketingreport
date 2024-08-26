// handlers/penjualan.go
package handlers

import (
	"database/sql"
	"net/http"

	"github.com/elhaqeeem/go-gin-mysql-marketingreport/models"
	"github.com/elhaqeeem/go-gin-mysql-marketingreport/utils"

	"github.com/gin-gonic/gin"
)

// CreatePenjualan creates a new penjualan record
func CreatePenjualan(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var p models.Penjualan
		if err := c.BindJSON(&p); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Generate a new transaction number
		transactionNumber, err := utils.GenerateTransactionNumber(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		p.TransactionNumber = transactionNumber

		// Automatically calculate GrandTotal
		p.GrandTotal = p.TotalBalance + p.CargoFee

		_, err = db.Exec(`
            INSERT INTO Penjualan (TransactionNumber, MarketingID, Date, CargoFee, TotalBalance, GrandTotal)
            VALUES (?, ?, ?, ?, ?, ?)`,
			p.TransactionNumber, p.MarketingID, p.Date, p.CargoFee, p.TotalBalance, p.GrandTotal,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}

// GetPenjualan retrieves a Penjualan record by ID
func GetPenjualan(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var penjualan models.Penjualan

		// Prepare the query
		row := db.QueryRow(`
            SELECT id, TransactionNumber, MarketingID, Date, CargoFee, TotalBalance, GrandTotal
            FROM Penjualan WHERE id = ?`, id)

		// Scan the result into the penjualan struct
		err := row.Scan(&penjualan.ID, &penjualan.TransactionNumber, &penjualan.MarketingID, &penjualan.Date, &penjualan.CargoFee, &penjualan.TotalBalance, &penjualan.GrandTotal)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, penjualan)
	}
}

// get all data
func GetallPenjualan(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Prepare the query to get all penjualan records
		rows, err := db.Query(`
            SELECT id, TransactionNumber, MarketingID, Date, CargoFee, TotalBalance, GrandTotal
            FROM Penjualan`)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close() // Ensure rows are closed after processing

		var penjualans []models.Penjualan // Slice to hold all penjualan records
		for rows.Next() {
			var penjualan models.Penjualan
			err := rows.Scan(&penjualan.ID, &penjualan.TransactionNumber, &penjualan.MarketingID, &penjualan.Date, &penjualan.CargoFee, &penjualan.TotalBalance, &penjualan.GrandTotal)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			penjualans = append(penjualans, penjualan) // Append each record to the slice
		}

		// Check for any error during iteration
		if err := rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return the list of penjualan records
		c.JSON(http.StatusOK, penjualans)
	}
}

// UpdatePenjualan updates an existing Penjualan record
func UpdatePenjualan(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var penjualan models.Penjualan
		if err := c.ShouldBindJSON(&penjualan); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		_, err := db.Exec(`
            UPDATE Penjualan
            SET  MarketingID = ?, date = ?, CargoFee = ?, TotalBalance = ?
            WHERE id = ?`,
			penjualan.TransactionNumber, penjualan.MarketingID, penjualan.Date, penjualan.CargoFee, penjualan.TotalBalance, penjualan.GrandTotal, penjualan.ID,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}

// DeletePenjualan deletes a Penjualan record by ID
func DeletePenjualan(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		_, err := db.Exec(`DELETE FROM Penjualan WHERE id = ?`, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "deleted"})
	}
}
