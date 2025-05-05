package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Group: /api
	api := r.Group("/api")
	{
		api.GET("/greet", func(c *gin.Context) {
			name := c.Query("name")
			if name == "" {
				name = "Guest"
			}
			c.JSON(http.StatusOK, gin.H{"message": "Hello, " + name + "!"})
		})

		api.GET("/users/:id", func(c *gin.Context) {
			id := c.Param("id")
			_, err := strconv.Atoi(id)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"user_id": id, "status": "User fetched"})
		})
	}

	/* ANOTHER WAY , WE CAN PASS THE HANDLER FUNCTIONS FOR EACH REQUEST RATHER THAN USING ANNONYMOUS FUNCTIONS LIKE WE DO IN JAVASCRIPT

						api := r.Group("/api")

				api.GET("/greet", greetHandler)
				api.GET("/users/:id", userHandler)

				func greetHandler(c *gin.Context) {
					name := c.Query("name")
					if name == "" {
						name = "Guest"
					}
					c.JSON(http.StatusOK, gin.H{"message": "Hello, " + name + "!"})
				}

				func userHandler(c *gin.Context) {
					id := c.Param("id")
					_, err := strconv.Atoi(id)
					if err != nil {
						c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
						return
					}
					c.JSON(http.StatusOK, gin.H{"user_id": id, "status": "User fetched"})
				}


	*/

	// Group: /admin
	admin := r.Group("/admin")
	{
		admin.GET("/dashboard", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Welcome to the admin dashboard!"})
		})
	}
}


/*
			Internally, gin.Default() returns a pointer to a gin.Engine.
			method signature for gin.Default() as we can clearly see it return pointer of *Engine and we name it as 'r' or 'router'
			So r is already of type *gin.Engine.
			func Default() *Engine --> method of gin *Engine is a return type

			That’s why we are accepting the type of *gin.Engine in routes.go:
			func RegisterRoutes(r *gin.Engine) 

			| Concept                     | Value           |
			| --------------------------- | --------------- |
			| `gin.Default()` returns     | `*gin.Engine`   |
			| `r` in `main.go` is of type | `*gin.Engine`   |
			| `RegisterRoutes` accepts    | `*gin.Engine`   |
			| ✅ So it's a perfect match   | No issue at all |


*/
