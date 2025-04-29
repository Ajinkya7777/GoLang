package main
import "fmt"

//named struct vs annonymous struct

type Author struct{
	name string
	branch string
	year int
}

type HR struct {
	details Author
}

type Student struct{
	name string
	branch string
	year int
}

type Teacher struct{
	name string 
	subject string
	exp int
	details Student
}

func main() {
	result := HR{
		details: Author{"John","Mech",2010},
	}

	fmt.Println("Details of Author")
	fmt.Println(result)

	result2 := Teacher{
		name:"Mike",
		subject:"Java",
		exp: 5,
		details: Student{"Jerry","IT",2},
	}

	fmt.Println("\nDetails of the Teacher")
	fmt.Println("Teacher's Name:",result2.name)
	fmt.Println("Subject Name:",result2.subject)
	fmt.Println("Exp in Years:",result2.exp)

	fmt.Println("\nDetails of the Student")
	fmt.Println("Student's Name:",result2.details.name)
	fmt.Println("Subject Branch:",result2.details.branch)
	fmt.Println("Year:",result2.details.year)

	

}