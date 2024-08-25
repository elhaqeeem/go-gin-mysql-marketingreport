// handlers/marketing.go
package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/elhaqeeem/go-gin-mysql-marketingreport/models"
	"github.com/elhaqeeem/go-gin-mysql-marketingreport/utils"

	"github.com/gin-gonic/gin"
)

// CreateMarketing handles POST requests to create a new marketing entry.
func CreateMarketing(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var marketing models.Marketing
		if err := c.ShouldBindJSON(&marketing); err != nil {
			utils.JSONResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		_, err := db.Exec("INSERT INTO marketing (name) VALUES (?)", marketing.Name)
		if err != nil {
			utils.JSONResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusCreated, marketing)
	}
}

func GetMarketing(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		// Step 1: Explain the query for debugging purposes
		rows, err := db.Query("EXPLAIN SELECT id, name FROM marketing WHERE id = ?", id)
		if err != nil {
			utils.JSONResponse(c, http.StatusInternalServerError, "Error explaining query")
			return
		}
		defer rows.Close()

		// Collect explain output (all rows and columns)
		var explainOutput []string
		columns, err := rows.Columns()
		if err != nil {
			utils.JSONResponse(c, http.StatusInternalServerError, "Error fetching columns")
			return
		}

		// Prepare a slice of interface{} to hold the values
		values := make([]interface{}, len(columns))
		for i := range values {
			var value []byte
			values[i] = &value
		}

		for rows.Next() {
			if err := rows.Scan(values...); err != nil {
				utils.JSONResponse(c, http.StatusInternalServerError, "Error reading explain data")
				return
			}

			var rowOutput string
			for i, val := range values {
				// Convert []byte to string
				rowOutput += fmt.Sprintf("%s: %s ", columns[i], string(*(val.(*[]byte))))
			}
			explainOutput = append(explainOutput, rowOutput)
		}

		// Optionally, log the explainOutput for analysis
		fmt.Println("Query Explain Output:", explainOutput)

		// Step 2: Execute the actual query to fetch the marketing entry
		var marketing models.Marketing
		queryStmt := "SELECT id, name FROM marketing WHERE id = ?"
		row := db.QueryRow(queryStmt, id)
		if err := row.Scan(&marketing.ID, &marketing.Name); err != nil {
			if err == sql.ErrNoRows {
				utils.JSONResponse(c, http.StatusNotFound, "Marketing not found")
			} else {
				utils.JSONResponse(c, http.StatusInternalServerError, "Error fetching data")
			}
			return
		}

		// Return the successful response
		c.JSON(http.StatusOK, marketing)
	}
}

func GetAllMarketing(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Step 1: Explain the query for debugging purposes
		rows, err := db.Query("EXPLAIN SELECT id, name FROM marketing")
		if err != nil {
			utils.JSONResponse(c, http.StatusInternalServerError, "Error explaining query")
			return
		}
		defer rows.Close()

		// Collect explain output (all rows and columns)
		var explainOutput []string
		columns, err := rows.Columns()
		if err != nil {
			utils.JSONResponse(c, http.StatusInternalServerError, "Error fetching columns")
			return
		}

		// Prepare a slice of interface{} to hold the values
		values := make([]interface{}, len(columns))
		for i := range values {
			var value []byte
			values[i] = &value
		}

		for rows.Next() {
			if err := rows.Scan(values...); err != nil {
				utils.JSONResponse(c, http.StatusInternalServerError, "Error reading explain data")
				return
			}

			var rowOutput string
			for i := range values {
				rowOutput += fmt.Sprintf("%s: %s ", columns[i], string(*(values[i].(*[]byte))))
			}
			explainOutput = append(explainOutput, rowOutput)
		}

		// Optionally, log the explainOutput for analysis
		fmt.Println("Query Explain Output:", explainOutput)

		// Step 2: Execute the actual query to fetch all marketing entries
		rows, err = db.Query("SELECT id, name FROM marketing")
		if err != nil {
			utils.JSONResponse(c, http.StatusInternalServerError, "Error fetching data")
			return
		}
		defer rows.Close()

		// Prepare a slice to hold all marketing entries
		var marketings []models.Marketing

		for rows.Next() {
			var marketing models.Marketing
			if err := rows.Scan(&marketing.ID, &marketing.Name); err != nil {
				utils.JSONResponse(c, http.StatusInternalServerError, "Error reading marketing data")
				return
			}
			marketings = append(marketings, marketing)
		}

		// Return the successful response with all marketing entries
		c.JSON(http.StatusOK, marketings)
	}
}

// UpdateMarketing handles PUT requests to update an existing marketing entry.
func UpdateMarketing(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var marketing models.Marketing
		if err := c.ShouldBindJSON(&marketing); err != nil {
			utils.JSONResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		_, err := db.Exec("UPDATE marketing SET name = ? WHERE id = ?", marketing.Name, id)
		if err != nil {
			utils.JSONResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, marketing)
	}
}

// DeleteMarketing handles DELETE requests to remove a marketing entry by ID.
func DeleteMarketing(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		_, err := db.Exec("DELETE FROM marketing WHERE id = ?", id)
		if err != nil {
			utils.JSONResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Marketing deleted"})
	}
}
