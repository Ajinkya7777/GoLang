package main
import "fmt"

type Address struct {
	Name string
	city string
	Pincode int
}

func main() {
	var a Address 
	fmt.Println(a) // by default all the values are set to their default values string will be empty and int will be 0

	// Declaring and initializing a struct using a struct literal
	a1 := Address{"Ajinkya","Latur",413512}
	fmt.Println("address1:",a1)

    // initializing a struct Naming fields while
	a2 := Address{Name: "Shubham",city:"Pune",Pincode:411011}
	fmt.Println("address2:",a2)

	// Uninitialized fields are set to their corresponding zero-value
	a3 := Address{Name: "Madrid"}
	fmt.Println("address3:",a3)

	//HOW TO ACCESS FIELDS OF A STRUCTS
	fmt.Println("City Name:",a1.city)

}