package main

import (
    "02_notes_app/config"
    "02_notes_app/models"
    "02_notes_app/routes"
    // "02_notes_app/controllers"
	"github.com/gin-contrib/cors"

    "github.com/gin-gonic/gin"
)

func main() {
    // Connect to the database
    config.InitDB()

    // Auto-migrate models (create tables if they don't exist)
    config.DB.AutoMigrate(&models.Note{}, &models.Category{})

    
    r := gin.Default()
	// r.Use(cors.Default())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Frontend React app URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"}, 
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true, 
	}))

    routes.RegisterNoteRoutes(r)

    
    r.Run(":8080")
}
