package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Custom middleware to log each request
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Process request
		c.Next()

		// Log format
		log.Printf("[%s] %s %s (%dms)",
			c.Request.Method,
			c.Request.URL.Path,
			c.ClientIP(),
			time.Since(start).Milliseconds(),
		)
	}
}


/*

Why is c.Next() Important?
c.Next() is crucial because it tells Gin to proceed to the actual handler for the request (the part that gives the user what they asked for, like fetching data or rendering a page).

Without c.Next(), the request would get stuck in the middleware, and the user would never get a response.

In Simple Terms:
c.Next() is like a green light that lets the next part of the process happen (the actual work of the request).

It’s saying, "Okay, I've done my job (logging), now let the main part of the request (like the route handler passes request to get , post , put , delete) do its thing."

After that, the middleware gets to check how long the request took and do any final work.


FLOW OF A GET REQUEST :

---

 1. Request Comes in

Let's say someone makes a GET request to the server, for example:


GET /greet?name=Ajinkya


 The browser sends this request to your web server (the one running the Gin app) on localhost:8080.

---

 2. Middleware Step (RequestLogger)

Once the request hits the server, the first thing that happens is the middleware gets triggered. In our case, the middleware is the RequestLogger() function.

Here's what happens in the middleware:

1. Log the Start Time:

    The RequestLogger middleware records the start time of the request.

   
   start := time.Now()
   

2. Process the Request:

    The middleware then calls c.Next(), which tells Gin to continue processing the request and move to the actual route handler (the part that handles the request, for example, r.GET("/greet", ...)).

   The request continues its journey to the next step (the route handler).

---

 3. Route Handler for /greet

The request now reaches the route handler for the /greet path.

Here’s the handler for /greet in your code:


r.GET("/greet", func(c gin.Context) {
    name := c.Query("name")  // Get the 'name' query parameter from the URL
    if name == "" {
        name = "Guest"  // If 'name' is empty, set it to "Guest"
    }
    c.JSON(http.StatusOK, gin.H{
        "message": "Hello, " + name + "!",
    })
})


 What happens here?

   The c.Query("name") gets the value of the query parameter name from the request URL (/greet?name=Ajinkya).
   If the parameter name exists, it uses that value (in this case, "Ajinkya").
   If name is missing (empty string), it defaults to "Guest".
   The server then responds with a JSON message:

    
    { "message": "Hello, Ajinkya!" }
    

---

 4. Middleware After the Request Handler (Post-Processing)

Once the route handler finishes and sends the response back to the user, the middleware gets another chance to run after the request is processed.

1. Log the Duration:

    The RequestLogger middleware now calculates the time it took to process the request (from the moment it started to the moment it finished).

   
   log.Printf("[%s] %s %s (%dms)",
       c.Request.Method,
       c.Request.URL.Path,
       c.ClientIP(),
       time.Since(start).Milliseconds(),
   )
   

    It logs the HTTP method (e.g., GET), the path of the request (/greet), the client IP (where the request came from), and the duration (in milliseconds).

---

 5. Response is Sent Back to the User

Finally, after logging the request details, the Gin framework sends the response back to the browser, which is:


{
    "message": "Hello, Ajinkya!"
}


This is what the user sees in the browser.

---

 Flow Recap:

1. User makes a GET request to /greet?name=Ajinkya.
2. RequestLogger middleware starts: logs the start time and calls c.Next() to let the request proceed.
3. Route handler (for /greet) handles the request:

    Fetches the name query parameter.
    Responds with a greeting message in JSON format ({ "message": "Hello, Ajinkya!" }).
4. RequestLogger middleware finishes: logs the request method, path, client IP, and time taken to process the request.
5. Response is sent back to the user with the greeting.

---

 Why Is This Flow Useful?

1. Middleware allows us to do something before and after the main work is done. In this case, it logs request details, which is very useful for debugging and performance monitoring.

2. The handler processes the request (i.e., fetches data, computes results, or performs actions) and sends the response.

3. The log provides valuable information like how long the request took, which is helpful for understanding the performance of your server.

---

 Visualizing the Flow:


[Request] --> [RequestLogger middleware] --> [Route Handler] --> [Response]
     ^                                       |
     |                                       v
   Logs start time                         Logs response time


---

*/