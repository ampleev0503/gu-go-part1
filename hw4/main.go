package main

import (
	"fmt"
	"sort"
)

// Задание № 1
type shape interface {
	perimeter() float64
}

type triangle struct {
	a float64
	b float64
	c float64
}

func (t triangle) perimeter() float64 {
	return t.a + t.b + t.c
}

type rectangle struct {
	a float64
	b float64
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.a + r.b)
}

func summPerimeter(shapes ...shape) (res float64) {

	for _, s := range shapes {
		res += s.perimeter()
	}
	return
}

// Задание № 2
type abonent struct {
	name  string
	phone int
}

type books []abonent

func (b books) Len() int           { return len(b) }
func (b books) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b books) Less(i, j int) bool { return b[i].name < b[j].name }

func main() {
	// Задание № 1
	t1 := triangle{3, 4, 5}
	fmt.Println("Периметр треугольника равен: ", t1.perimeter())

	r1 := rectangle{3, 4}
	fmt.Println("Периметр прямоугольника равен: ", r1.perimeter())

	fmt.Println("Сумма периметров: ", summPerimeter(t1, r1))
	fmt.Println("_______________________________________________")

	// Задание № 2
	abonents := []abonent{
		{"Vadim", 89998887777},
		{"Aleksandr", 89998887776},
		{"Andrey", 89998887775},
		{"Kirill", 89998887774},
	}

	sort.Sort(books(abonents))

	fmt.Println(abonents)

}
