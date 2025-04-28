package main
import "fmt"

func main() {
	var a uint8 = 225;
	fmt.Println(a,a-3)

	var b int16 = 32767
	fmt.Println(b+2,b-2)



// Possible arithmetic operations for intigers
	var x int16 = 10;
	var y int16 = 20

	//Addition
	fmt.Printf("Addition: %d + %d = %d\n", x,y,x+y)
	//Subtraction
	fmt.Printf("Subtraction: %d - %d = %d\n", x,y,x-y)
	//Multiplication
	fmt.Printf("Multiplication: %d * %d = %d\n", x,y,x*y)
	//Division
	fmt.Printf("Division: %d / %d = %d\n", x,y,x/y)
	//Modulus
	fmt.Printf("Remainder: %d %% %d = %d\n", x,y,x%y)

	//local declration of a variable (Float Example)
	fmt.Printf("\n")
	c := 20.45
	fmt.Println(c)
	fmt.Printf("The type of c is : %T\n", c)
	//We can perform same operations on float data types as well except the Moduls operator

	
	// Complex Numbers

	var comp1 complex128 = complex(6, 2)
	var comp2 complex64  = complex(9, 2)

	fmt.Println(comp1)
	fmt.Println(comp2)
	fmt.Printf("\nThe type of comp1 is %T & "+"The Type of comp2 is %T",comp1,comp2)
	// get real part
    realNum := real(comp1)
    fmt.Println("\nReal part of complex number 1:", realNum)
    // get imaginary part
    imaginary := imag(comp2)
    fmt.Println("Imaginary part of complex number 2:", imaginary)


	//BOOLEANS

	str1 := "Ajinkya"
	str2 := "ajinkya"
	str3 := "Ajinkya"

	result1:= str1 == str2
	result2:= str1 == str3

	fmt.Println( result1) //false
    fmt.Println( result2) //true


fmt.Printf("The type of result1 is %T and "+"the type of result2 is %T",result1, result2)

	//STRING -> Immutable

	str := "Ajinkya"

	fmt.Printf("\nLength of the sring is: %d",len(str))

	fmt.Printf("\nString is: %s",str)

	fmt.Printf("\n Type of str is: %T", str)

	// String concatenation

	firstName := "Ajinkya"
	lastName  := " Shinde"

	fmt.Println("\nNew String :",firstName + lastName);

}