package main

import (
	"log"
	"pratikum-repository-skripsi/config"
	"pratikum-repository-skripsi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Koneksi ke database
	config.ConnectDB()

	// Inisialisasi router Gin
	router := gin.Default()

	// Setup Routes (panggil folder routes)
	routes.SetupRoutes(router)

	// Running
  if err := router.Run(":8080"); err != nil {
        log.Fatal("Server gagal dijalankan:", err)
    }


}