package main

import "fmt"

func main15() {
	// fmt.Println("learning pointers")

	// what is a pointer?
	// ?=  variable that stores the memory address of another varibale

	// declaring a pointer?
	// *T => type of the pointer variable which points to the value of T

	// b := 900
	// var a *int = &b
	// fmt.Println("address of b is ", a)
	// fmt.Printf("type of a is %T ", a)
	// b := "thousand"
	// var c *string = &b
	// fmt.Println(c)
	// fmt.Printf("the value of b %s", *c)

	// the & is used to get the address of the variable

	// zero value of a pointer
	// one := 78
	// var two *int
	// if two == nil {
	// 	fmt.Println("two is ", two)
	// 	two = &one
	// 	fmt.Println("two after initialization ", two)
	// }

	fmt.Println()

	// creating pointers using the new  function
	size := new(int)
	fmt.Printf("size is %d, Type is %T, value is %v\n", *size, size, size)
	*size = 90
	fmt.Println("value is now ", *size)

	fmt.Println()

	//dereferencing a pointer
	// accessing the value of the variable which the pointer points to
	a := "hello, world"
	b := &a
	fmt.Println("address of a is ", b)
	fmt.Println("value of a is ", *b)

	fmt.Println()

	one := "javascript is awesome"
	two := &one
	fmt.Println("address of one is ", two)
	fmt.Println("value of one is ", *two)
	*two = "golang is awesome"
	fmt.Println("address of one is ", two)
	fmt.Println("value of one is ", *two)

	fmt.Println()

	aa := 78
	fmt.Println("value of aa bofore is ", aa)
	bb := &aa
	change(bb)
	fmt.Println("value of aa after is ", aa)

	fmt.Println()

	// returning a pointer from a function
	d := returnPointer()
	fmt.Println("value of d: ", *d)

	// go does not support pointer arithmethic

}

// passing a pointer to a function
func change(value *int) {
	*value = 100
}

// return a pointer from func
func returnPointer() *int {
	one := 45
	return &one
}
