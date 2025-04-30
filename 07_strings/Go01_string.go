package main
import "fmt"

func main(){
	//there are two ways to create string 
	 //1.using double quotes

	 	//using var
	 	var myString string
		myString = "Hello"

		fmt.Println("MyString:",myString)

		//using short hand :=
		 myString2 := "Hello World"
		 fmt.Println("My String2:",myString2)

	 //2.using string literals i.e. backticks
	    
	 myString3 := `New Hello World....!`
	 fmt.Println("My String3:",myString3)


	 //length of a string
	 fmt.Println("MyString Length:",len(myString))
	 fmt.Println("MyString2 Length:",len(myString2))
	 fmt.Println("MyString3 Length:",len(myString3))

	 //how to iterate over a string using range for loop
	 for index, char := range myString {
		fmt.Printf("The char at %d is %c\n",index,char)
	 }

}