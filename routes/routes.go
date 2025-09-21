package routes

import (
	"pratikum-repository-skripsi/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	skripsi := router.Group("/skripsi")
	{
		skripsi.POST("/", controllers.CreateSkripsi)     // Create
		skripsi.GET("/", controllers.GetAllSkripsi)      // Get All
		skripsi.GET("/:id", controllers.GetSkripsiByID)  // Get by ID
		skripsi.PATCH("/:id", controllers.PatchSkripsi)  // Update Datas
		skripsi.DELETE("/:id", controllers.DeleteSkripsi) // Delete
	}
}

