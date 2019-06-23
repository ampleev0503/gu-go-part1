package main

import (
	"fmt"
)

func main() {

	fmt.Print("____________________________\n\n")

	var number int
	fmt.Print("Введите число: ")
	fmt.Scanln(&number)
	fmt.Printf("число %d является ", number)
	if isDividedBy(number, 2) {
		fmt.Println("четным")
	} else {
		fmt.Println("нечетным")
	}
	fmt.Scanln()

	fmt.Print("____________________________\n\n")

	fmt.Print("Введите число: ")
	fmt.Scanln(&number)
	fmt.Printf("число %d делится на 3 ", number)
	if isDividedBy(number, 3) {
		fmt.Println("без остатка")
	} else {
		fmt.Println("с остатком")
	}
	fmt.Scanln()

	fmt.Print("____________________________\n\n")
	fmt.Println("100 первых чисел Фибоначчи: ")
	fmt.Print("0 ")
	fibonacci(0, 1, 1)
}

func isDividedBy(number, divider int) bool {
	if number % divider == 0 {
		return true
	}

	return false
}

func fibonacci(prev, curr, count int) {
	if count == 100 {
		return
	}
	fmt.Printf("%d ", prev + curr)
	count++ 
	fibonacci(curr, prev + curr, count)
}
