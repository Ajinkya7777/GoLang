package main
import "fmt"

func sum(nums ...int)int{
	total := 0

	for _,n :=range nums {
		total += n
	}
	return total
}

func greet(name string,nums ...int){
	fmt.Println(name)
	for _,n := range nums{
		fmt.Println("Numbers:",n)
	}
}

func main() {
	fmt.Println("Sum of 1,2,3:",sum(1,2,3))
	fmt.Println("Sum of 4,5:",sum(4,5))
	fmt.Println("Sum of no numbers:",sum())
fmt.Printf("============================\n")
//Variadic Functions with Other Parameters
	greet("Sum of numbers:", 1, 2, 3)
    greet("Another sum:", 4, 5)
    greet("No numbers sum:")

}