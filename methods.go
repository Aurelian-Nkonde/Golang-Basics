package main

import "fmt"

type Employee struct {
	name   string
	salary int
}

func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %d ", e.name, e.salary)
}

type Student struct {
	name string
	age  int
}

func (s Student) studentProfile() {
	fmt.Printf("student name is %s, and %d years ", s.name, s.age)
}

func main() {
	// fmt.Println("methods in go")

	// what is a methos?
	// just a function with a special reciver type between the func keyword and the method name

	/*
		func (t Type) methodName(parameter list){

		}
	*/
	// emply1 := Employee{
	// name:   "jason",
	// salary: 900,
	// }
	// emply1.displaySalary()

	oneStudent := Student{
		name: "jerome",
		age:  13,
	}
	oneStudent.studentProfile()
}
