/* 
Pointer to a Struct :- A pointer to a struct allows you to directly reference and modify the data in the original struct without making a copy.
	*) Declaring a Pointer to a Struct : There are two common ways to create a pointer to a struct in Go.
	   1) Using the & operator to get the memory address of an existing struct.
	   2) Using the new function, which returns a pointer to a newly allocated struct.
    *) Accessing Struct Fields Through a Pointer
		In Go, we don’t need to explicitly dereference the pointer (using *) to access the fields. 
		Go automatically dereferences pointers when accessing fields, so personPointer.name works directly without (*personPointer).name.
*/


package main
import "fmt"


//define struct 
type Person struct {
	name string
	age int
}


func main() {
	//instance of a struct
	p1 := Person{ name:"A",age:25}

	//creating pointer to struct
	//1) Using the & operator to get the memory address of an existing struct.
	personPointer := &p1

	//accessing the field using the pointer
	fmt.Println("Name:",personPointer.name) // Automatically dereferences

	//modifying struct using pointer
	personPointer.age = 26
	fmt.Println("Updated struct:",p1)


	//creating pointer to a new instance of Person
	// 2) Using the new function, which returns a pointer to a newly allocated struct.
	personePointer2 := new(Person)
	personePointer2.name = "B"
	personePointer2.age = 30

	fmt.Println("New Person:",*personePointer2)
}

