//how to pass pointers to function 

package main
import "fmt"


func passPointers(a *int){
	//dereferencing -> p will be 7
	*a = 7
}
func main(){
	p := 14
	a := &p 
	fmt.Println("Before changing p:",p)
	passPointers(a);   //passing pointer
	// passPointers(&p); //passing address 
	fmt.Println("After changing p:",p)

	//There are two ways to pass the pointers
	 //1) Passing an address of the variable 	-> &p
	 //2) Passing the pointer variable directly -> a

}