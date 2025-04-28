package main

import "fmt"

func main() {

	var v1 int = 100
	var v2 int = 70

	//normal if-else
	if (v1 < 1000) {
		fmt.Printf("v1 is less than 1000\n")

		//nested if-else
		if(v2 == 70){
			fmt.Printf("v2 is equal to 700\n")
		}else{
			fmt.Printf("v2 is not equal to 700\n")
		}

	}else{
		fmt.Printf("v1 is greater than 1000\n")
	}

	//if..else..if ladder

	var a int = 300
	if (a == 100){
		fmt.Printf("value of a is 100\n")
	}else if(a == 200){
		fmt.Printf("value of a is 200\n")
	}else if(a == 300){
		fmt.Printf("value of a is 300\n")
	}else{
		fmt.Printf("None of the values is matching..\n")
	}

}