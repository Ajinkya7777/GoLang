package main

import "fmt"

type Book struct {
    Title  string
    Author string
}

func updateBookDetails(b *Book) {
    b.Title = "New Title"  // This changes the ORIGINAL book
}

func main() {
    book := Book{Title: "Old Title", Author: "John"}
    
    updateBookDetails(&book)  // Passing a POINTER (reference) to the book
    
    // Now the original 'book' is modified, because we passed a reference
    fmt.Println(book.Title)  // Outputs: "New Title"
}


/*

Passing by pointer allows Gin to modify the original struct with the data from the incoming request body.

*/