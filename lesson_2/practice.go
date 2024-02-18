package main

import (
	"fmt"
	"math"
)

/*
Задачи по переменным и типам данных:

    Создайте программу, которая объявляет переменные различных типов данных (int, float64, string) и выводит их значения.

    Напишите программу, которая принимает два числа от пользователя, складывает их и выводит результат.

    Создайте функцию, которая принимает на вход число и возвращает его квадрат.

    Напишите программу, которая использует короткий синтаксис для объявления переменных и выводит их значения.
*/

func square(number int) float64 {
	return math.Pow(float64(number), float64(2))
}

func main() {
	// 1.
	var age int = 20
	_ = age

	var price float64 = 123.321
	_ = price

	var name string = "Dastan"
	_ = name

	fmt.Println(age)
	fmt.Println(price)
	fmt.Println(name)

	fmt.Println()

	fmt.Printf("Type: %T\n", age)
	fmt.Printf("Type: %T\n", price)
	fmt.Printf("Type: %T\n", name)

	//var word string

	//fmt.Print("Your word: ")
	//_, _ = fmt.Scan(&word) // Так мы сохраняем полученный объект в переменную
	//fmt.Printf("Your word - %v\n", word)

	// 2.
	/*
		var a, b int

		fmt.Print("Input your numbers: ")
		fmt.Scan(&a, &b) // Так мы сохраняем полученный объект в переменную
		fmt.Println(a + b)
	*/

	// 3.
	fmt.Println(square(5))
}
