package main
import "fmt"

func main(){
	// Anonymous function
	func() {
		fmt.Println("hello world...!")
	}()

	//Assigning to a Variable	
	func2 := func() {
		fmt.Println("hello world...!")
	}
	func2()

	//Passing Arguments
	func(name string){
		fmt.Println("hello",name)
	}("world as args...!")

	//Passing Anonymous func as Arguments to another func
	func3 := func(p,q string) string {
		return p + q + "...!"
	}
	helloworld(func3)

	// Returning Anonymous Functions
		func4 := helloworld2()
		fmt.Println(func4("hello ","world"))

}

func helloworld(i func(p,q string) string){
	fmt.Println(i("hello"," world"))
}
func helloworld2() func(p,q string) string{
	myf := func(p,q string) string {
		return p + q + "...!"
	}
	return myf
}