// Defer :- In Go language, defer statements delay the execution of the function or method or an anonymous method until the nearby functions returns
//1. In Go language, multiple defer statements are allowed in the same program and they are executed in LIFO
package main
import "fmt"

func mul(a1, a2 int) int {
 
    res := a1 * a2
    fmt.Println("Result: ", res)
    return 0
}
 
func show() {
    fmt.Println("this will print before defer func")
}

func add(a1, a2 int) int {
    res := a1 + a2
    fmt.Println("Result: ", res)
    return 0
}
 
// Main function
func main() {
 
    // Calling mul() function
    // Here mul function behaves
    // like a normal function
    mul(23, 45)
 
    // Calling mul()function
    // Using defer keyword
    // Here the mul() function
    // is defer function
    defer mul(23, 56)
 
    // Calling show() function
    show()

fmt.Printf("another example \n")
	fmt.Println("Start")
 
    // Multiple defer statements
    // Executes in LIFO order
    defer fmt.Println("End")
    defer add(34, 56)
    defer add(10, 10)
}