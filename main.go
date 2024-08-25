package main

import (
	"log"

	"github.com/elhaqeeem/go-gin-mysql-marketingreport/config" // Import config package correctly
	"github.com/elhaqeeem/go-gin-mysql-marketingreport/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Initialize the database
	gin.SetMode(gin.ReleaseMode)
	db := config.InitDB()
	defer db.Close()

	// Set up Gin router
	r := gin.Default()

	// Define routes here
	// Marketing CRUD routes
	r.POST("/marketing", handlers.CreateMarketing(db))
	r.GET("/marketing/:id", handlers.GetMarketing(db))
	r.GET("/marketing", handlers.GetAllMarketing(db))
	r.PUT("/marketing/:id", handlers.UpdateMarketing(db))
	r.DELETE("/marketing/:id", handlers.DeleteMarketing(db))
	// CRUD routes for Penjualan
	r.POST("/penjualan", handlers.CreatePenjualan(db))
	r.GET("/penjualan/:id", handlers.GetPenjualan(db))
	r.GET("/penjualan", handlers.GetallPenjualan(db))
	r.PUT("/penjualan/:id", handlers.UpdatePenjualan(db))
	r.DELETE("/penjualan/:id", handlers.DeletePenjualan(db))
	// get data
	r.GET("/komisi", handlers.GetKomisi(db))
	r.POST("/pembayaran", handlers.CreatePembayaran(db))
	r.GET("/pembayaran", handlers.GetPembayaran(db))
	r.GET("/angsuran/:pembayaran_id", handlers.GetAllAngsuran(db))
	r.GET("/angsuran/status/:pembayaran_id", handlers.CheckInstallmentStatus(db))
	// Start server
	r.Run(":8080")
}
