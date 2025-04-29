package main
import "fmt"


func main() {

	//array of 5 integers, all initalized to 0
	// var nums[5]int
	
	arr1 := [5]int{50,2,3,4,5} //full initialization
	arr2 := [...]int{10,20,30} //Compiler infers the size(lenght of an array)
	
	//accessing elements
	fmt.Println("array1[0]",arr1[0])
	//modify element
	arr2[1] = 100
	
	for i :=0;i<len(arr1);i++ {
		fmt.Println(arr1[i])
	}

	fmt.Printf("\n")

	//range for loop
	for i, value := range arr2 {
    fmt.Printf("Index: %d, Value: %d\n", i, value)
	}
}


/*
Important Observations About Array

In an array, if an array does not initialize explicitly, then the default value of this array is 0
In an array, you can find the length of the array using len() method 
*/