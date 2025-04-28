package main
import "fmt"

//Pointers:- Pointers in Go programming language or Golang is a variable which is used to store the memory address of another variable. You can also create a pointer to a struct

type Employee struct {
	firstName,lastName string
	age,salary int
}

func main() {
	// taking pointer to struct
	emp := &Employee{"John","Doe",35,8000}
	
	fmt.Println("address",emp)
	fmt.Println("value",*emp)
	fmt.Println("firstname:",emp.firstName) //*emp.firstName we can directly use emp.firstName shorthand
	fmt.Println("lastname:",emp.lastName)
}