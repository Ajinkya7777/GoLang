package main

import "fmt"
// Go doesn't use parentheses around conditions in for loops.
func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
fmt.Printf("===============\n")
//var is not valid inside the for loop initializer , With var declared before the loop

var a int
for a = 5; a < 10; a++ {
	fmt.Println(a)
}

// Infinite loop 
    // for { 
    //   fmt.Printf("i can do it...\n")   
    // } 

//for loop as while Loop
d:= 0 
    for d < 3 { 
       d += 2 
    } 
  fmt.Println("d =",d)

//enhanced for-each loop -> similar like java int num : nums
numbers := []int{10, 20, 30, 40}

//without index
for _, number := range numbers {
    fmt.Println(number)
}

//with index
for index, number := range numbers {
    fmt.Println(index,number)
}


//map
fmt.Printf("map example\n")
//map declaration
scores := map[string]int{
    "Alice": 90,
    "Bob":   85,
}

for name, score := range scores {
    fmt.Printf("%s: %d\n", name, score)
}


}
