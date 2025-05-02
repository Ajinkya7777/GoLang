/* 
Pointers :- Pointers in Go programming language or Golang is a variable that is used to store the memory address of another variable.
			 Pointers in Golang is also termed as the special variables.

			 Below are two operators used with pointers:-
	1) * Operator also termed as the dereferencing operator used to declare pointer variable and access the value stored in the address.
    2) & operator termed as address operator used to returns the address of a variable or to access the address of a variable to a pointer.

	IMP points while working with pointer: 
	1) If you are specifying the data type along with the pointer declaration then the pointer will be able to handle the memory address of that specified data type variable. 
		For example, if you taking a pointer of string type then the address of the variable that you will give to a pointer will be only of string data type variable, not any other type.
		example : 
		var s *string = &a -> here 'a' variable should be of type string or else we will get error
    2) To overcome the above mention problem you can use the Type Inference concept of the var keyword. There is no need to specify the data type during the declaration. 
	   The type of a pointer variable can also be determined by the compiler like a normal variable. Here we will not use the * operator. 
		It will internally determine by the compiler as we are initializing the variable with the address of another variable.
		example : 
		var y = 5
		var p = &y -> pointer type will be internally determined by the compier while initializing the pointer.
	3)  You can also use the shorthand (:=) syntax to declare and initialize the pointer variables. 
		The compiler will internally determine the variable is a pointer variable if we are passing the address of the variable using &(address) operator to it.
		 example :
		 y := 5
		 p := &y
*/

package main
import "fmt"

func main(){
	var x int = 5748

	//declaration of pointer -> default value or zero-value of pointer is always nil.
	var p *int 
	fmt.Println("default value of pointer: ",p)
	//initialization of pointer -> address will be in hexadecimal
	p = &x

	//declaration & initialization of the pointers can be done into a single line.
	 var q *int = &x
	fmt.Println("value of q pointer :",q)

	fmt.Println("value stored in x = ",x)
	fmt.Println("Address of x = ",&x)
	fmt.Println("value stored in variable p = ",p)

	//Dereferencing the Pointer :-  As we know that * operator is also termed as the dereferencing operator. It is not only used to declare the pointer variable but also used to access the value stored in the variable which the pointer points to which is generally termed as indirecting or dereferencing. * operator is also termed as the value at the address of.
	fmt.Println("value stored in x(*p) = ",*p)

	//Changing the value of the pointer of at the memory location instead of assigning a new value to the variable
	fmt.Println("value stored in x(*p) before changing = ",*p)
	*p = 5
	fmt.Println("value stored in x(*p) after changing = ",*p)
	fmt.Println("value stored in x(*p) after changing = ",x)
}