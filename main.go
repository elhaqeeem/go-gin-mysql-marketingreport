package main

import (
	"log"
	"os/exec"

	"github.com/elhaqeeem/go-gin-mysql-marketingreport/config"
	"github.com/elhaqeeem/go-gin-mysql-marketingreport/handlers"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	gin.SetMode(gin.ReleaseMode)
	db := config.InitDB()
	defer db.Close()

	// Start React app using npm
	cmd := exec.Command("npm", "start")
	cmd.Dir = "./frontend" // Change directory to the frontend folder
	if err := cmd.Start(); err != nil {
		log.Fatalf("Failed to start React app: %v", err)
	}
	log.Println("React app started")

	// Set up Gin router
	r := gin.Default()
	r.Use(cors.Default())

	// Define routes here
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

	// Get data routes
	r.GET("/komisi", handlers.GetKomisi(db))
	r.POST("/pembayaran", handlers.CreatePembayaran(db))
	r.GET("/pembayaran", handlers.GetPembayaran(db))
	r.GET("/angsuran/:pembayaran_id", handlers.GetAllAngsuran(db))
	r.GET("/angsuran/status/:pembayaran_id", handlers.CheckInstallmentStatus(db))

	// Serve static files and the React app
	r.Static("/static", "./frontend/build/static")
	r.StaticFile("/manifest.json", "./frontend/build/manifest.json")
	r.StaticFile("/favicon.ico", "./frontend/build/favicon.ico")
	r.StaticFile("/logo192.png", "./frontend/build/logo192.png")
	r.StaticFile("/logo512.png", "./frontend/build/logo512.png")
	r.StaticFile("/robots.txt", "./frontend/build/robots.txt")

	// Serve the React app's index.html at the root
	r.NoRoute(func(c *gin.Context) {
		c.File("./frontend/build/index.html")
	})

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
