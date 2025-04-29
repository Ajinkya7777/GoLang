package main
import "fmt"

func main(){
	var original = [5]int{10,20,30,40,50}
    var copy [5]int

	//using for loop 
	for i:=0;i<len(original);i++{
		copy[i] = original[i]
	}

	fmt.Println("Original:",original)
	fmt.Println("Copy:",copy)

	//using direct assignment, does not work with slices
	var copy2 [5]int = original
	fmt.Println("Copy2:",copy2)

	//using pointers
	var copy3 *[5]int = &original
	fmt.Println("Copy3:",*copy3)
}