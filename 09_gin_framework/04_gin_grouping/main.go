package main

import (
	"04_gin_grouping/middleware"
	"04_gin_grouping/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Global middleware
	r.Use(middleware.RequestLogger())

	// Register grouped routes
	routes.RegisterRoutes(r)

	r.Run(":8080")
}
