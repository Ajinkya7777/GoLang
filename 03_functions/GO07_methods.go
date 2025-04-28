//Go methods are like functions but with a key difference: they have a receiver argument, which allows access to the receiver’s properties. The receiver can be a struct or non-struct type, but both must be in the same package

package main 
import "fmt"

// Defining a struct
type person struct{
	name string
	age int
}

// Creating a custom type based on int
type number int


// Defining a method with struct receiver -> Methods with Struct Type Receiver
func (p person) display(){
	fmt.Println("Name:",p.name)
	fmt.Println("Age:",p.age)
}

// Defining a method with a non-struct receiver
func (n number) square() number {
    return n * n
}

func main(){
	// Creating an instance of the struct
	a := person{name:"Ajinkya",age:25}

	a.display()


	// Creating an instance of the struct
	 c := number(4)
    d := c.square()

    fmt.Println("Square of", c, "is", d)
}