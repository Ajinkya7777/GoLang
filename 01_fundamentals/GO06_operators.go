package main
import "fmt"


func main() {

	p:= 20
	q:= 10

//Arithmetic -> (+ - * / %)
	//+
		result1 := p+q
		fmt.Printf("Result of p + q = %d\n",result1)
	//-
		result2 := p-q
		fmt.Printf("Result of p - q = %d\n",result2)
	//*
		result3 := p*q
		fmt.Printf("Result of p * q = %d\n",result3)
	//'/'
		result4 := p/q
		fmt.Printf("Result of p / q = %d\n",result4)
	//%
		result5 := p%q
		fmt.Printf("Result of p %% q = %d\n",result5)




//Relational
		// ‘=='(Equal To) 
   result11:= p == q 
   fmt.Println(result11) 
      
   // ‘!='(Not Equal To) 
   result12:= p != q 
   fmt.Println(result12) 
      
   // ‘<‘(Less Than) 
   result13:= p < q 
   fmt.Println(result13) 
      
   // ‘>'(Greater Than) 
   result14:= p > q 
   fmt.Println(result14) 
      
   // ‘>='(Greater Than Equal To) 
   result15:= p >= q 
   fmt.Println(result15) 
      
   // ‘<='(Less Than Equal To) 
   result16:= p <= q 
   fmt.Println(result16)
   
//Logical

	if(p!=q && p>=q){  
        fmt.Println("True") 
    } 
        
    if(p!=q || p<=q){  
        fmt.Println("True") 
    } 
        
    if(!(p==q)){  
        fmt.Println("True") 
    } 


//Bitwise
//Assignment
//Misc Operators

var z int = 4

   r := &z
   fmt.Println(*r)

   *r = 7

   fmt.Println(z)



}



//Arithmetic -> (+ - * / %)


//Relational
//Logical
//Bitwise
//Assignment
//Misc