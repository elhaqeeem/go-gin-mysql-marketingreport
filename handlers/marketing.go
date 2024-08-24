// handlers/marketing.go
package handlers

import (
	"database/sql"
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

// GetMarketing handles GET requests to fetch a marketing entry by ID.
func GetMarketing(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var marketing models.Marketing
		row := db.QueryRow("SELECT id, name FROM marketing WHERE id = ?", id)
		if err := row.Scan(&marketing.ID, &marketing.Name); err != nil {
			if err == sql.ErrNoRows {
				utils.JSONResponse(c, http.StatusNotFound, "Marketing not found")
			} else {
				utils.JSONResponse(c, http.StatusInternalServerError, err.Error())
			}
			return
		}

		c.JSON(http.StatusOK, marketing)
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
