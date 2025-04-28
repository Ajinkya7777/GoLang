package main
import "fmt"

func main() {
	//calling func
	result := multiply(2,3)
	fmt.Printf("multiplication: %d\n",result)

	//call by value
	x := 5
	y := 5

	fmt.Printf("Before: x = %d, y = %d\n",x,y)
	res := multi(x,y)
	fmt.Printf("multiplication: %d\n",res)
	fmt.Printf("After: x = %d, y = %d\n",x,y)

fmt.Printf("========================\n")

	//call by reference
	p := 5
	q := 5

	fmt.Printf("Before: p = %d, q = %d\n",p,q)
	res2 := multi2(&p,&q)
	fmt.Printf("multiplication: %d\n",res2)
	fmt.Printf("After: p = %d, q = %d\n",p,q)

}

// name(args) returntype
func multiply(a,b int) int{
	return a * b
}



//call by value
func multi(a,b int) int{
	a = a * 2
	return a * b
}


//call by reference
func multi2(a,b *int) int{
	*a = *a * 2
	return *a * *b
}