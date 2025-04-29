package main
import "fmt"

type Student struct{
	details struct {
		name string
		age int
	}
	subject string
}

//Anonymous Fields
type Employee struct{
	string //employee name  
	int    //employee salary
}

func main() {
	student := Student{
		details: struct {
			name string
			age int
		}{
			name: "John",
			age: 25,
		},
		subject:"Java",
	}

	fmt.Println("Name:",student.details.name)
	fmt.Println("Age:",student.details.age)
	fmt.Println("Subject:",student.subject)

	employee := Employee{"Jacob",25000}
	fmt.Println("\nEmployee Details")
	fmt.Println("Employee Name:",employee.string)
	fmt.Println("Employee Salary:",employee.int)
}

/*
Important Points about Anonymous Fields
1. Uniqueness Requirement: You cannot use two fields of the same type in a struct. 
		type InvalidStudent struct {
			int
			int // Error: duplicate type
		}
2. Combining Named and Anonymous Fields: You can mix anonymous and named fields within a struct.
		type Student struct {
			id int // Named field
			int    // Anonymous field
		}

*/