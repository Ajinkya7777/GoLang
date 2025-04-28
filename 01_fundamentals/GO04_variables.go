package main

import "fmt"

func main() {
	//var variable_name type = expression
	//NOTE: In the above syntax, either type or =expression can be omitted, but not both.

	//var myVar1 = 20; // =20 is expression
	var myVar1 int //variable declared & initialized w/o expression

	//var myVar2 = "Ajinkya"
	var myVar2 string 
	fmt.Printf("The value of myVar1 is: %d\n",myVar1)
	fmt.Printf("The type of myVar1 is: %T\n",myVar1)
	fmt.Printf("The value of myVar2 is: %s\n",myVar2)
	fmt.Printf("The type of myVar2 is: %T\n",myVar2)

	//NOTE 1 : if u use type, then u are allowed to declare multiple variables of the same type in the single declaration

	var var1 , var2 , var3 int = 2,3,4

	fmt.Printf("The values are %d %d %d",var1,var2,var3)

	//NOTE 2 : if u remove type, then u are allowed to declare multiple variables of a different type in the single declaration. The type of variables is determined by the initialized values.

	var v1 , v2 = 2 , "AJ"

	fmt.Printf("\nThe values are %d %s",v1,v2)

	//NOTE 3 : Short Variable Declaration
	   // while working with short variable declaration the type of the variable is determined by the type of the expression so in below example price is type of int
        //variable_name := expression -> price := 20
		// := ----> declaration
		// =  ----> assignment

		myShortVar1 := 39
		myShortVar2 := "Ajinkya"

		fmt.Printf("\nThe value of myShortVar1 is: %d",myShortVar1)
		fmt.Printf("\nThe type of myShortVar1 is: %T",myShortVar1)
		fmt.Printf("\nThe value of myShortVar2 is: %s",myShortVar2)
		fmt.Printf("\nThe type of myShortVar2 is: %T",myShortVar2)


		// Using short variable declaration 
		// Multiple variables of different types 
		// are declared and initialized in the single line 
		myvar11, myvar22, myvar33 := 800, "Geeks", 47.56

 }