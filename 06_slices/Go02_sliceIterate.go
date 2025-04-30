package main
import "fmt"

func main(){
	fruits := []string{"Apple","Banana","Kiwi","Dragon"}

	//for loop
	for i:= 0; i<len(fruits);i++{
		fmt.Println(fruits[i])
	}

	//range for in loop
	for index,fruit := range fruits {
		fmt.Printf("Index = %d and fruit = %s\n",index,fruit)
	}

	//blank identifier in for loop
	for _,fruit := range fruits{
		fmt.Println("Fruit:",fruit)
	}
}

/*
Modifying Slice: As we already know that slice is a reference type it can refer an underlying array. So if we change some elements in the slice, then the changes should also take place in the referenced array. Or in other words, if you made any changes in the slice, then it will also reflect in the array
Comparison of Slice: In Slice, you can only use == operator to check the given slice is nill or not. If you try to compare two slices with the help of == operator then it will give you an error
*/