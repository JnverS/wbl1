package main

import (
	"fmt"
	"math"
)

/*
	Разработать программу нахождения расстояния между двумя точками, которые
	представлены в виде структуры Point с инкапсулированными параметрами x,y и
	конструктором.
*/

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

//расстояние между двумя точками AB = √(xb - xa)2 + (yb - ya)2
func Distance(a *Point, b *Point) float64 {
	return math.Sqrt(math.Pow((b.x-a.x), 2)+math.Pow((b.y-a.y), 2))
}

func main() {
	a := NewPoint(3, 4)
	b := NewPoint(2, 3)

	fmt.Println(Distance(a, b))
}
