package handlers

import (
	"database/sql"
	"net/http"

	"github.com/elhaqeeem/go-gin-mysql-marketingreport/models"
	"github.com/elhaqeeem/go-gin-mysql-marketingreport/utils"

	"github.com/gin-gonic/gin"
)

// GetKomisi returns the commission for each marketing.
func GetKomisi(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query(`
            SELECT marketing_id, DATE_FORMAT(date, '%Y-%m') AS bulan, SUM(grand_total) AS omzet
            FROM Penjualan
            GROUP BY marketing_id, DATE_FORMAT(date, '%Y-%m')
        `)
		if err != nil {
			utils.JSONResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		defer rows.Close()

		var results []models.Komisi

		for rows.Next() {
			var k models.Komisi
			err := rows.Scan(&k.MarketingId, &k.Bulan, &k.Omzet)
			if err != nil {
				utils.JSONResponse(c, http.StatusInternalServerError, err.Error())
				return
			}
			if k.Omzet >= 500000000 {
				k.KomisiPersen = 10
			} else if k.Omzet >= 200000000 {
				k.KomisiPersen = 5
			} else if k.Omzet >= 100000000 {
				k.KomisiPersen = 2.5
			} else {
				k.KomisiPersen = 0
			}
			k.KomisiNominal = (k.Omzet * k.KomisiPersen) / 100
			results = append(results, k)
		}

		c.JSON(http.StatusOK, results)
	}
}
