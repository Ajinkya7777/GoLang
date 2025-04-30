package main
import (
    "fmt"
    "sort"
)

func main() {
    intSlice := []int{42, 23, 16, 15, 8, 4}
	slc1 := []string{"Python", "Java", "C#", "Go", "Ruby"}
    
	 fmt.Println("Before sorting:")
    fmt.Println("Slice 1: ", slc1)
	fmt.Println("intSlice: ", intSlice)
   

	// Sorting in ascending order
	sort.Strings(slc1)
	sort.Ints(intSlice)


	fmt.Println("\nAfter sorting:")
    fmt.Println("Slice 1: ", slc1)
    fmt.Println("intSlice: ", intSlice)
}