package main

import (
	"fmt"
	"unsafe"
)

var text string

//const name = "Dastan"

func main() {
	//var word string
	//word = "Hello"

	//word := "Hello, World"
	//fmt.Printf("Type: %T\nValue: %v", word, word)

	//isStudent := true
	//fmt.Printf("\n%T", isStudent)

	//text = "Apple"
	//fmt.Println(text)

	//fmt.Println(name)

	//age := "20"
	//sprint := fmt.Sprintf("I'm age: %s", age)
	//_ = sprint

	//fmt.Println(sprint)

	//var num1 int64 = 15
	//var num2 int = 15

	//fmt.Println(num1 + int64(num2))

	fmt.Println(unsafe.Sizeof(uint8(1)))
	fmt.Println(unsafe.Sizeof(uint64(1)))

	fmt.Println(text)
}
