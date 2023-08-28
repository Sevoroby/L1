package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func (p Point) CalculateDistance(second Point) float64 {
	return math.Sqrt(((second.x - p.x) * (second.x - p.x)) + ((second.y - p.y) * (second.y - p.y)))
}

func main() {
	// Создание объектов точек
	point1 := NewPoint(5, 0)
	point2 := NewPoint(7, 4)

	fmt.Printf("Точка 1: x=%.2f,y=%.2f\n", point1.x, point1.y)
	fmt.Printf("Точка 2: x=%.2f,y=%.2f\n", point2.x, point2.y)

	// Расчёт расстояния между двумя точками
	distance := point1.CalculateDistance(point2)
	fmt.Printf("Расстояние между точками: %.5f\n", distance)
}
