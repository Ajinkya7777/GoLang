//blank indentifier :  - The real use of Blank Identifier comes when a function returns multiple values, but we need only a few values and want to discard some values. 
package main
import "fmt"

func main() {
	mul, _ := mul_div(105,7)

	//only using the mul variable
	fmt.Println("105 x 7 = ", mul)
}

//function returning two values of integer type
func mul_div(n1 int,n2 int) (int , int){
	return n1 * n2, n1/n2
}