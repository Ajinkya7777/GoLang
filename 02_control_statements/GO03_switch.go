package main
import "fmt"

func main(){

	day := 3
	days := 3

	//  Basic switch 
	//No need for break — Go breaks automatically after each case
	//default is optional, just like in Java
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Another day")
	}

	//Multiple values in one case:
	switch days {
	case 1, 2, 3:
		fmt.Println("Weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Unknown day")
}


//with interface
    describe(42)
	describe("hello")
	describe(true)
}

//with interfaces
func describe(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("It's an int: %d\n", v)
	case string:
		fmt.Printf("It's a string: %s\n", v)
	case bool:
		fmt.Printf("It's a bool: %t\n", v)
	default:
		fmt.Println("Unknown type")
	}
}