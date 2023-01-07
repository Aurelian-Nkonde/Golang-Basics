package main

import "fmt"

func main11() {
	// holds values of same time eg 1,4,5,7
	//declaration => [n]T, n is number/length, T denotes the type

	var a [3]int
	fmt.Println(a)

	//all elements are asigned to zero value of the element type
	//index starts at 0 and ends at length - 1
	a[0] = 905
	a[1] = 673
	a[2] = 900
	fmt.Println(a)

	b := [2]string{}
	b[0] = "chibaba"
	b[1] = "maita basa"
	fmt.Println(b)

	c := [...]int{1, 3, 56, 89} //letting the compiler determine/figure out the length
	fmt.Println(len(c))

	// the size of the array is part of the array type
	famA := [...]string{"mom", "papa", "boi"}
	famB := famA
	famB[0] = "mai"
	fmt.Println(famA)
	fmt.Println(famB)

	//getting the length of an array
	d := [...]int{9, 4, 32, 45, 6, 7, 92}
	fmt.Println(len(d))

	// using range looping
	numbs := [3]int{5, 7, 2}
	sum := 0
	for _, v := range numbs {
		sum += v
	}
	fmt.Println("the sum of ", numbs, " is ", sum)

	//range with both key and value
	// for i,v := range {}

	//ignoring the key
	// for _, v := range {}

	//ignoring tha value
	// for i := range {}

	// one := [2][3]int{
	// 	{3, 4, 5},
	// 	{7, 8, 9},
	// }

	// array2(one)

	var two [2][2]string
	two[0][0] = "oi"
	two[0][1] = "ma"
	two[1][0] = "gaga"
	two[1][1] = "dhakwaz"

	// for _, v := range two {
	// 	for _, v1 := range v {
	// 		fmt.Println(v1)
	// 	}
	// 	fmt.Println()
	// }

	// SLICES
	// a slice is represented by []T
	//creating a slice from an array
	aa := [4]int{3, 5, 12, 3}
	bb := aa[0:3] //[start:end]
	fmt.Println(bb)

	// a slice does not own data
	// dd := [7]int{6, 7, 3, 4, 2, 8, 9}
	// fmt.Println("before ", dd)
	// ee := dd[0:2]
	// for i := range ee {
	// 	ee[i]++
	// }
	// fmt.Println("after ", dd)

	// length and capacity of a slice
	xx := []int{3, 5, 6}
	fmt.Println(len(xx), " is the length, and capacity is ", cap(xx))

	//syntax => make([]T, length, capacity), capacity part is optional and defaults to 0
	//creating a slice using make
	xxx := make([]string, 2, 4)
	xxx[0] = "thousand"
	xxx[1] = "mathaza"
	fmt.Println(xxx)

	//APPEND TO A SLICE
	//arrays are restricted to a  fixed length, the length cannot be increased
	cars := make([]string, 1)
	cars[0] = "tesla"
	cars = append(cars, "toyota")
	cars = append(cars, "mazda")
	fmt.Println(cars)

	// the zero value of a slice type is nil

	// var nils []string
	// if nils == nil {
	// 	nils = append(nils, "hola senorita")
	// }
	// fmt.Println(nils)

	fruits := []string{"mango", "apple"}
	animals := []string{"dog", "cat", "lion"}
	livingThings := append(fruits, animals...)
	fmt.Println(livingThings)

	// muiltidimensional slices
	oi := [][]string{
		{"java", "c#"},
		{"c", "c++"},
		{"javascript", "python"},
	}
	for _, v := range oi {
		for _, v1 := range v {
			fmt.Println(v1)
		}
		fmt.Printf("\n")
	}
}
