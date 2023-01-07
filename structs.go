package main

import "fmt"

func main16() {
	// fmt.Println("structs")

	// what is a struct?
	// is a user defined type that represents a collection of fields

	// defining a struct
	type Doctor struct {
		firstName string
		gender    string
		age       int
	} //defines struct type Doctor with fields firstName, gender and age
	// the above is called =>  a named struct

	type Nurse struct {
		name      string
		age       int
		hostpital string
	}

	// first choice syntax
	anna := Nurse{ //creating struct specifying field names
		name:      "anna-maria",
		age:       35,
		hostpital: "st luis",
	}
	fmt.Println("Nurse 1", anna)

	maggie := Nurse{"magret", 40, "st james"} //creating struct type without specifying fields names
	fmt.Println("Nurse 2 ", maggie)

	fmt.Println()

	// creating anonymous  structs!
	// it is possible to create structs without declaring new data type
	// are called anonymous types
	coder1 := struct {
		name     string
		age      int
		language string
	}{
		name:     "thousand",
		age:      90,
		language: "c#, go",
	}

	fmt.Println(coder1)
	fmt.Println()

	// accessing individual elements of a struct type
	type Teacher struct {
		name  string
		grade int
	}
	jon := Teacher{
		name:  "mr jones",
		grade: 2,
	}
	fmt.Println("Teacher name is: ", jon.name)
	fmt.Println("Teacher teaches grades: ", jon.grade)
	fmt.Println()

	// pointers to a struct
	type Car struct {
		name  string
		color string
		year  int
	}
	firstCar := &Car{
		name:  "toyota",
		color: "red",
		year:  2019,
	}
	fmt.Println("name of car is ", (*firstCar).name)
	fmt.Println("year made of car is ", firstCar.year)

	fmt.Println()

	// nested structs
	// struct contains a field which is itself a struct

	type Address struct {
		city  string
		state string
	}

	type Person struct {
		name    string
		age     int
		address Address
	}

	thaza := Person{
		name: "thousand",
		age:  90,
		address: Address{
			city:  "chicago",
			state: "western Cape",
		},
	}
	fmt.Println(thaza)
	fmt.Println("user name is ", thaza.name)
	fmt.Println("user lives in ", thaza.address.city)
}
