package utils

import "github.com/gin-gonic/gin"

// JSONResponse is a utility function to send JSON responses.
func JSONResponse(c *gin.Context, statusCode int, message interface{}) {
	c.JSON(statusCode, gin.H{"error": message})
}
