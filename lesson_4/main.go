package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Hello, world")

	// Функция как значение (callback)
	var multiplier func(a, b int) int
	multiplier = func(a, b int) int { return a * b }
	_ = multiplier
	//fmt.Println(multiplier(2, 2))

	divider := func(a, b int) float64 { return float64(a) / float64(b) }
	_ = divider
	//fmt.Println(divider(5, 2))

	// Приём функции как аргументо
	sumFunc := func(a, b int) int { return a + b }
	subtractFunc := func(a, b int) int { return a * b }

	_, _, _ = sumFunc, subtractFunc, calculate

	//fmt.Println(calculate(10, 10, sumFunc))
	//fmt.Println(calculate(10, 5, subtractFunc))

	// Возврат функции из функции
	divideBy2 := createDivider(2)
	divideBy5 := createDivider(5)

	fmt.Println(divideBy2(100))
	fmt.Println(divideBy5(100))

	// Нейминг как модификатор доступа
	// Если мы хотим использовать из другого пакета мы назваем его с большой буквы и если нет то малелькой

	// Замыкание
	dollar := 20
	getDollarValue := func() int {
		return dollar
	}

	fmt.Println(getDollarValue())

	dollar = 50
	fmt.Println(getDollarValue())
}

func calculate(a, b int, action func(a, b int) int) int {
	return action(a, b)
}

func createDivider(divider int) func(a int) int {
	dividerFunc := func(a int) int {
		return a / divider
	}
	return dividerFunc
}
