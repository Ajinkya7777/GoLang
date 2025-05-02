/* 
     Go Pointer to Pointer (Double Pointer):- 
			A pointer is a special variable so it can point to a variable of any type even to a pointer. Basically, this looks like a chain of pointers. 
			When we define a pointer to pointer then the first pointer is used to store the address of the second pointer.
			This concept is sometimes termed as Double Pointers.
*/


package main
import "fmt"

func main(){
	//variable of type int 
	var v int = 100

	//taking pointer to store addrs of 'v'
	var p1 *int = &v

	//taking pointer to store addrs of pointer p1
	var p2 **int = &p1

	fmt.Println("The value of variable v = ",v)
	fmt.Println("Address of variable v = ",&v)

	fmt.Println("The value of p1 = ",p1)
	fmt.Println("Address of p1 = ",&p1)

	fmt.Println("The value of p2 is = ",p2)

	//Dereferencing the pointer to pointer 
    fmt.Println("Value at the address of p2 is or *p2 = ", *p2)

	//double pointer will give the value of variable V 
	fmt.Println("*(Value at the address of p2 is) or **p2 = ", **p2) 

}