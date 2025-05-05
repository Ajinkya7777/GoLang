package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv" 
)
  /*
  packages info...
1) github.com/gin-gonic/gin -> This is the Gin web framework, which helps build web applications in Go, similar to Spring Boot for Java. It provides routing, request handling, and response formatting
2) strconv -> string to integer conversion
3)  net/http -> This package provides constants for HTTP status codes like http.StatusOK (200), http.StatusNotFound (404) 
  */
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: 1, Title: "Go Basics", Author: "John"},
	{ID: 2, Title: "Gin Framework", Author: "Jane"},
}

func main() {
	r := gin.Default()

	// GET all books
	r.GET("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, books)
	})

	/*
	GET single book by ID
	 1) /books/:id -> id is a path parameter
	 2) c.Param("id") gets the ID from the URL (like /books/1), and strconv.Atoi() converts the string ID to an integer (since URL parameters are always strings).
	*/
	r.GET("/books/:id", func(c *gin.Context) {

		/*
		 The following line converts the "id" path parameter from string to int using strconv.Atoi().
		 If the URL parameter is not a valid number (e.g., "/books/abc"), this will result in an error.
		 In this case, we're ignoring the error (with _), assuming the "id" parameter is always valid.
		 In production code, it is recommended to check the error and return a bad request response if invalid.
		*/
		id, _ := strconv.Atoi(c.Param("id"))

		/*
			Loop through the list of books to find a match for the given ID.
			We don't need the index value, so we use '_' to discard it, and 'b' is the book object at each iteration.
		*/
		for _, b := range books {
			if b.ID == id {
				c.JSON(http.StatusOK, b)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
	})

	/* 
	POST - Add new book
	 1) c.BindJSON(&newBook): This is similar to @RequestBody in Spring Boot. 
	 						  It binds the JSON body of the request to the newBook object. It is a method from the Gin framework 
							  that tries to bind the JSON data from the incoming request body to the updatedBook struct.
	 2) General Syntax of for range Loop
	 	
	 		//syntax
	 	for index, value := range collection {
    		// Use 'index' and 'value' here
		}

			//our code
		for _, b := range books {
    		// Use 'b' as the book object here
		}

	 3) General Syntax
	   id, err := strconv.Atoi(c.Param("id"))

	   //we are ignoring the error by using the _ (underscore) to discard it 
	   id, _ := strconv.Atoi(c.Param("id"))

	*/
	r.POST("/books", func(c *gin.Context) {
		var newBook Book
		if err := c.BindJSON(&newBook); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newBook.ID = len(books) + 1
		books = append(books, newBook)
		c.JSON(http.StatusCreated, newBook)
	})

	// PUT - Update book by ID
	r.PUT("/books/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var updatedBook Book
		if err := c.BindJSON(&updatedBook); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for i, b := range books {
			if b.ID == id {
				books[i].Title = updatedBook.Title
				books[i].Author = updatedBook.Author
				c.JSON(http.StatusOK, books[i])
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
	})

	// DELETE - Remove book by ID
	r.DELETE("/books/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		for i, b := range books {
			if b.ID == id {
				books = append(books[:i], books[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
	})

	r.Run(":8080")
}


/*

Q1) Why we have a such a weird syntax in terms of handling error? 

		This is a feature in Go that may seem a bit unusual if you're coming from a Java or other more traditional object-oriented language. The syntax you're referring to is a combination of declaration and error checking in a single if statement.

		In Go, you can declare variables inside the if condition, and this is a very idiomatic and powerful feature of the language. It allows for a cleaner and more compact way to both declare and check errors in one step.

		Let's break it down more clearly:

		The if err := ...; err != nil syntax explained

		This is what you're asking about:


		if err := c.BindJSON(&updatedBook); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}


		err := c.BindJSON(&updatedBook):

		Here, you're calling c.BindJSON(&updatedBook) which tries to bind the incoming JSON body to the updatedBook object.
		c.BindJSON() returns two values:

			1. A result (in this case, it’s the successfully bound struct, but we’re ignoring it here).
			2. An error value, which is what we are interested in here.
		err is assigned the error result of c.BindJSON(&updatedBook). If there is no error, err will be nil.

		if err != nil:

		This checks whether the error (err) is not nil. If it's not nil, it means there was an error in binding the JSON (e.g., the request body is malformed or the JSON structure doesn’t match the updatedBook struct).
		If err != nil, the code inside the if block will be executed — in this case, returning a 400 Bad Request response with the error message.

		Why is the syntax written this way?

		1. Cleaner Code:
		Go allows you to declare and check errors inline, which reduces unnecessary lines of code. You don’t need a separate error declaration and an if check like in other languages.

		For example, in Java or other languages, you might do something like this:

		java code for the same
		Error err = someFunction();
		if (err != null) {
			// handle error
		}
		

		But in Go, this can be written as:

		go
		if err := someFunction(); err != nil {
			// handle error
		}
		

		2. Scoped Variable:
		The variable err is only scoped to the if block. This means that err is automatically declared and initialized in the if statement, and it's not accessible outside the if block. This helps avoid unnecessary global variables or cluttering up the function with a variable that’s only used once.

		For example:

		go
		if err := c.BindJSON(&updatedBook); err != nil {
			// err is scoped here
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 'err' is not accessible here
		

		3. Idiomatic Go:
		This is a common Go idiom. Go emphasizes simplicity and readability. Using this pattern, the if block both declares the variable and checks for an error in one go, making the code concise and easy to read.

		Alternative Syntax: Declaring err Outside

		If you're more comfortable with declaring variables separately, you can still declare err outside of the if statement and use it in a more traditional if block. Here's how it would look:


		var err error
		err = c.BindJSON(&updatedBook)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}


		In this case, we explicitly declare err first and then assign the result of c.BindJSON(&updatedBook) to it. Afterward, we check the err variable.
		This works the same way but takes a few more lines and doesn't have the scoped benefit of the first example.

		Why use the if err := ... pattern?

		Concise: The code is more compact and avoids unnecessary variable declarations.
		Less boilerplate: You don't have to declare the error variable before using it.
		Scoped error: The err variable is scoped only to the if block, which is safer and avoids polluting the broader scope.
		More Go idiomatic: This pattern is widely used in the Go community because it makes code easier to read and understand without needing extra boilerplate.

		Example in Context

		r.PUT("/books/:id", func(c gin.Context) {
			// Convert the "id" URL parameter to an integer
			id, _ := strconv.Atoi(c.Param("id"))
			
			// Declare and bind the JSON body to 'updatedBook'
			if err := c.BindJSON(&updatedBook); err != nil {
				// If the binding fails, return a 400 Bad Request with the error message
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			// Search for the book with the given ID and update it
			for i, b := range books {
				if b.ID == id {
					books[i].Title = updatedBook.Title
					books[i].Author = updatedBook.Author
					c.JSON(http.StatusOK, books[i])
					return
				}
			}

			// If the book was not found, return a 404 Not Found
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		})


		In this handler:

		if err := c.BindJSON(&updatedBook); err != nil: We attempt to bind the JSON to the updatedBook struct. If the binding fails, we return an error.
		The error variable err is scoped only inside the if statement and is not available elsewhere in the function.

		Summary

		if err := ...; err != nil is a Go idiom to declare and check errors in a single statement. It's concise, clean, and idiomatic in Go.
		It allows the error to be scoped to just the if block, keeping the code short and readable.
		It's similar to declaring the error outside the if statement but avoids boilerplate code.

 Q2) Why Does BindJSON Need a Pointer?

		The Core Reason: Mutability and Modification of Data
		In Go, Gin’s BindJSON method needs to modify the struct that it’s binding the incoming JSON data into.

		By passing a value (e.g., newBook): If you passed the value of the newBook struct to BindJSON, it would only modify a copy of that struct inside the function, not the original one. So, after the function call, the changes wouldn’t be reflected in the actual newBook variable you declared in your handler.

		By passing a pointer (e.g., &newBook): When you pass a pointer to the BindJSON method, Gin can directly modify the original newBook struct that you declared. This is critical because you want BindJSON to update the struct with the data coming from the request body

		Why You Always Pass a Pointer to BindJSON
		BindJSON essentially populates the struct fields (e.g., Title, Author) with the incoming JSON data. If you pass a pointer, it can modify the original struct directly.

		If you pass the struct by value (without &), Gin cannot modify your original struct, because it would only be modifying a local copy inside the BindJSON function. So, the original struct you passed to BindJSON would remain unchanged.

