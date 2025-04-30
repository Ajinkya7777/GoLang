package main
import "fmt"

func main() {

	//syntax for slice
	//var myslice[]T
	//var myslice[]T{}
	//var myslice[]T{v1,v2,v3}
	
	array := [5]int{1,2,3,4,5}
	slice := array[1:4]

	fmt.Println("Array:",array)
	fmt.Println("Slice:",slice)

	slice1 := []int{1,2,3}
	slice1 = append(slice1,4,5,6)
	fmt.Println("Slice1:",slice1)

	arr := [7]string{"This", "is", "the", "tutorial", "of", "Go", "language"}
	fmt.Println("Array2:",arr)

	myslice := arr[1:6]

	fmt.Println("MySlice:",myslice)

	fmt.Println("Slice Length:",len(myslice))
	fmt.Println("Slice Capacity:",cap(myslice))


	//How to create and initialize a Slice?

	//1.Using Slice Literal:
	var mySlice1 = []string{"hello","world"}
	fmt.Println("My Slice 1:",mySlice1)

	mySlice2 := []int{1,2,3,4,5}
	fmt.Println("My Slice 2:",mySlice2)

	//2.Using an Array
	fruits :=[4]string{"Apple","Banana","Dragon","Kiwi"}

	  var fruitSlice1 = fruits[1:2]
	  fruitSlice2 := fruits[0:] //[0:fruit.length]
	  fruitSlice3 := fruits[:2] //[0:2] default 0
	  fruitSlice4 := fruits[:]  //[0:fruits.length]

	  fmt.Println("Fruits:",fruits)
	  fmt.Println("FruitSlice1:",fruitSlice1)
	  fmt.Println("FruitSlice2:",fruitSlice2)
	  fmt.Println("FruitSlice3:",fruitSlice3)
	  fmt.Println("FruitSlice4:",fruitSlice4)

	//3. Using already existing slice
	 originalSlice := []int{10,20,30,40,50,60,70}

	 var my_Slice1 = originalSlice[1:5]
	 my_Slice2 := originalSlice[0:]
	 my_Slice3 := originalSlice[:6]
	 my_Slice4 := originalSlice[:]
	 my_Slice5 := originalSlice[2:4]

	 fmt.Println("Original Slice:",originalSlice)
	 fmt.Println("New Slice1:",my_Slice1)
	 fmt.Println("New Slice2:",my_Slice2)
	 fmt.Println("New Slice3:",my_Slice3)
	 fmt.Println("New Slice4:",my_Slice4)
	 fmt.Println("New Slice5:",my_Slice5)

    //4. Using make() function: 

}