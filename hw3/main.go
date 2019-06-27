package main

import (
	"fmt"
	"geek-university/gu-go-part1/hw3/queue"
	"geek-university/gu-go-part1/hw3/structs"
)



func main() {

	car1 := structs.Car{
		Brand:"audi",
		Year: 2007, 
		TrunkCapacity: 400, 
		LaunchEngine: false, 
		OpenWindows: false, 
		HowFullTrunk: 30,
	}

	car2 := structs.Car{
		Year: 2008,
		Brand: "bmw",
		TrunkCapacity: 410,
	}

	car1.LaunchEngine = true

	fmt.Println("Автомобиль 1: ", car1)
	fmt.Println("Автомобиль 2: ", car2)
	fmt.Println("_________________________")

	queue.Push("Меня")
	queue.Push("Зовут")
	queue.Push("Вадим")

	fmt.Println(queue.GetQueue())

	queue.Pop()
	fmt.Println(queue.GetQueue())
	queue.Pop()
	fmt.Println(queue.GetQueue())
	queue.Pop()
	fmt.Println(queue.GetQueue())

	fmt.Println(queue.Pop())



}