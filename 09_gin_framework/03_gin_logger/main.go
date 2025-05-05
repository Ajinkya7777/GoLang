package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Apply custom logger middleware
	r.Use(RequestLogger())

	// 1. Query parameter: /greet?name=Ajinkya
	r.GET("/greet", func(c *gin.Context) {
		name := c.Query("name")
		if name == "" {
			name = "Guest"
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, " + name + "!",
		})
	})

	// 2. Path parameter: /users/42
	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")

		// Optional: convert string ID to int
		_, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid user ID",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user_id": id,
			"status":  "User fetched successfully",
		})
	})

	
	r.Run(":8080")
}
