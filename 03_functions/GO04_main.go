//main and init functions 

package main
import ("fmt" "sort" "strings" "time")

func main(){
	fmt.Println("Welcome to main() function")
	s := []int {5,4,3,2,1}
	sort.Ints(s);
	fmt.Println("Sorted :",s)
}

// Multiple init() function
func init() {
    fmt.Println("Welcome to init() function")
}

func init() {
    fmt.Println("Hello! init() function")
}


//INIT Functions : - 
//init() function is just like the main function, does not take any argument nor return anything. 
// This function is present in every package and this function is called when the package is initialized