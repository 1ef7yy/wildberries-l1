package main

import (
	"fmt"
	"math"
)

// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с
// инкапсулированными параметрами x,y и конструктором.

type Point struct {
	x int
	y int
}

func NewPoint(x, y int) Point {
	return Point{
		x: x,
		y: y,
	}
}

func FindDistance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(float64(p1.x-p2.x), 2) + math.Pow(float64(p1.y-p2.y), 2))
}

func main() {
	p1 := NewPoint(0, 0)
	p2 := NewPoint(3, 4)

	distance := FindDistance(p1, p2)

	fmt.Printf("Distance between points (%d, %d) and (%d, %d) is %f\n", p1.x, p1.y, p2.x, p2.y, distance)
}
