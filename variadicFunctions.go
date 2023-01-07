package main

func main12() {
	// what is a variadic function?
	// => a func that accepts a variable of arguments, ellipsis [...]

	// syntax
	// func a(a int, c ...int){}

	// one(2, 9, 4, 5)
	// variadic can only be last

	// Canfound(3, 5, 7, 32, 54, 3)
}

/*
func Canfound(num int, nums ...int) {
	fmt.Printf("type of nums is %T ", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println("num found at index ", i)
			found = true
		}
	}
	if !found {
		fmt.Println("num is not found in nums")
	}
	fmt.Printf("\n")
}
*/
