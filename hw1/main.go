package main

import (
	"fmt"
	"math"
	"strconv"
)

const ConvDollarToRouble = 65

func main() {
	fmt.Println("Конвертация рублей в доллары:")
	fmt.Println("____________________________")

	var roubles float64
	fmt.Println("Сколько рублей Вы желаете перевести в доллары: ")
	fmt.Scanln(&roubles)

	fmt.Println("Вы получите " + strconv.FormatFloat(covertRoublesToDollar(roubles), 'f', 2, 64) + " долларов")
	fmt.Println("\nДля продолжения нажмите enter")
	fmt.Scanln()

	fmt.Println("Нахождение площади, периметра и гипотенузы по катетам:")
	fmt.Println("______________________________________________________")

	var a, b float64
	fmt.Print("Чему равен катет a: ")
	fmt.Scanln(&a)
	fmt.Print("Чему равен катет b: ")
	fmt.Scanln(&b)

	c := findHypotenuse(a, b)
	fmt.Println("Гипотенуза равна " + strconv.FormatFloat(c, 'f', 2, 64))
	fmt.Println("Площадь равна " + strconv.FormatFloat(findSquare(a, b), 'f', 2, 64))
	fmt.Println("Периметр равен " + strconv.FormatFloat(findPerimeter(a, b, c), 'f', 2, 64))

	fmt.Println("\nДля продолжения нажмите enter")
	fmt.Scanln()

}

func covertRoublesToDollar(roubles float64) float64 {
	return roubles / ConvDollarToRouble
}

func findHypotenuse(a, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func findSquare(a, b float64) float64 {
	return a * b / 2
}

func findPerimeter(a, b, c float64) float64 {
	return a + b + c
}
