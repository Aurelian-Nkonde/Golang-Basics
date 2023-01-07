package main

import "fmt"

func main14() {
	// strings
	// a string => is a slice of byte in go
	// name := "thousand wema tricks"
	// fmt.Println(name)

	// name := "thousand mathaza"
	// fmt.Printf("string: %s\n", name)
	// printCharacters(name)

	// something := "nyikaIno"
	// fmt.Println(something)
	// printChars1Runes(something)

	// RUNE
	// a builtln type in go, its the aliasis of the int32

	// string concanation
	// name1 := "thousand"
	// name2 := "wematricks"
	// fullname := name1 + " " + name2
	// fmt.Println(fullname)

	// strings are immutable => once created they cannot be changed
}

// func printBytes(word string) {
// 	fmt.Printf("bytes: ")
// 	for i := 0; i < len(word); i++ {
// 		fmt.Printf("%x\n ", word[i])
// 	}
// }

// func printCharacters(word string) {
// 	fmt.Println("characters: ")
// 	for i := 0; i < len(word); i++ {
// 		fmt.Printf("%c ", word[i])
// 	}
// }

func printChars1(word string) {
	fmt.Println("characters: ")
	for i := 0; i < len(word); i++ {
		fmt.Printf("%c ", word[i])
	}
}

func printChars1Runes(word string) {
	fmt.Println("characters: ")
	runes := []rune(word)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}
