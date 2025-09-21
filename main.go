package main

import (
	"log"
	"pratikum-repository-skripsi/config"
	"pratikum-repository-skripsi/routes"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Koneksi ke database
	config.ConnectDB()

	// Inisialisasi router Gin
	router := gin.Default()

	// Konfigurasi CORS custom
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // hanya izinkan frontend React (URL FE)
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Setup Routes (panggil folder routes)
	routes.SetupRoutes(router)

	// Running server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
