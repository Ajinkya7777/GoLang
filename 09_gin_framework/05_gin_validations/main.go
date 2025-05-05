package main

import (
	"05_gin_validations/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("/users", func(c *gin.Context) {
		var user models.User

		// Bind and validate request
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// If validation passed
		c.JSON(http.StatusOK, gin.H{
			"message": "User created successfully",
			"user":    user,
		})
	})

	r.Run(":8080")
}
