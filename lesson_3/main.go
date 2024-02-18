package main

import "fmt"

/*
func funcName(params...) (returned values){
	function's body
}
*/

func main() {
	//Greet()

	//name := "Dastan"
	//PersonGreet(name)

	//name, surname := "Dastan", "Kubatbaew"
	//BioGreet(name, surname)

	//sum := Sum(1, 2)
	//fmt.Println(sum)

	//sum, multi := MultiSum(2, 34)
	//fmt.Println(sum, multi)

	a, b := nameMultiSum(2, 4)
	fmt.Println(a, b)
}

func Greet() {
	fmt.Println("Hello")
}

func PersonGreet(name string) {
	fmt.Printf("Hello, %s", name)
}

func BioGreet(name, surname string) {
	fmt.Printf("Hello, %s %s", name, surname)
}

func Sum(a, b int) int {
	return a + b
}

func MultiSum(a, b int) (int64, int64) {
	return int64(a + b), int64(a) * int64(b)
}

func nameMultiSum(a, b int) (sum int, multi int) {
	sum = a + b
	multi = a * b
	return
}
