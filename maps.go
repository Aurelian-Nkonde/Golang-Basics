package main

import "fmt"

func main13() {
	//MAPS!
	//? => a builtIn type that stored key-value pairs
	// similar to python dictionaries

	// a := map[int]string{
	// 1: "java",
	// 2: "c#",
	//} // creating a map
	// fmt.Println(a)

	// employees := make(map[string]string)

	// var b = make(map[int]string)
	// b[1] = "toyota"
	// fmt.Println(b)

	countries := make(map[int]string)
	fmt.Println(countries)
	// adding stuff to a map
	countries[0] = "usa"
	countries[1] = "zambia"
	countries[2] = "jamaica"
	fmt.Println(countries)

	// initializing a map during the declaration
	var players = map[int]string{
		1: "calonga",
		2: "master",
	}
	fmt.Println(players)

	// retrieving single value from a map
	fmt.Println(players[1], " is the first player from the list")
	// the zero value of a map is nil

	// checking if key exist in a map
	//  value, ok := map[key]
	// wise := map[string]string{
	// 	"thaza": "the goat of go",
	// 	"king":  "the king of norway",
	// }
	// value, ok := wise["thaza"]
	// if ok == true {
	// 	fmt.Println("this king ", value, " exists!!")
	// 	fmt.Println("long hail king")
	// } else {
	// 	fmt.Println("the king ", value, " is long dead")
	// }

	// iterating over elements of a map
	// cities := map[int]string{
	// 	1: "oslo",
	// 	2: "maddison",
	// 	3: "london",
	// 	4: "berlin",
	// }
	// for _, v := range cities {
	// 	fmt.Println(v, " is a good city")
	// }

	// deleting items in a map  ==> delete(map, key)
	// fmt.Println(cities)
	// delete(cities, 2)
	// fmt.Println(cities)

	// map of structs

	person1 := person{
		age:      90,
		language: "shona",
	}

	person2 := person{
		age:      24,
		language: "xhosa",
	}
	coders := map[string]person{
		"thaza": person1,
		"jay":   person2,
	}
	for name, info := range coders {
		fmt.Println(name, " info => ", info)
	}

}

type person struct {
	age      int
	language string
}
