package main

import "fmt"

func main() {
	//var firstVariable bool
	//fmt.Println(firstVariable)

	//secondVariable := true
	//fmt.Println(secondVariable)

	age := 15

	// Базовый синтаксис
	if age < 18 {
		fmt.Println("Ты молод")
	}

	// Короткий синтаксис
	if isChild := isChildren(age); isChild == true {
		fmt.Println("Молодой")
		fmt.Println(isChild)
	}

	// Если ... Иначе
	age = 19
	if age > 18 {
		fmt.Println("Ты старый")
	} else {
		fmt.Println("Ты молод")
	}

	// && - И, and
	age = 10
	if age >= 6 && age <= 18 {
		fmt.Println("Ты школьник")
	}

	// || - или, or
	age = 40
	if age == 20 || age == 40 {
		fmt.Println("Возраст 20 или 40")
	}

	// ! - не, not
	if !isChildren(age) {
		fmt.Println("Ты старый")
	}
}

func isChildren(age int) bool {
	return age < 18
}
