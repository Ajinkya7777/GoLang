/*
Understanding Pointers vs Values in Go

Value: When you pass a variable by value (e.g., just newBook), you're passing a copy of the variable to the function. Any changes made to that copy won’t affect the original variable.

Pointer: When you pass a variable by reference (using &newBook), you're passing the address (pointer) to the original variable. This means any changes made to the variable inside the function will directly affect the original variable.

Passing by value would mean Gin can’t modify the original struct, because it’s just working with a copy of the struct.
*/

package main

import "fmt"

type Book struct {
    Title  string
    Author string
}

func updateBookDetails(b Book) {
    b.Title = "New Title"  // This changes ONLY the COPY of the book
}

func main() {
    book := Book{Title: "Old Title", Author: "John"}
    
    updateBookDetails(book)  // Passing a copy (by value)
    
    // The original 'book' is NOT modified, because we passed a copy
    fmt.Println(book.Title)  // Outputs: "Old Title"
}


