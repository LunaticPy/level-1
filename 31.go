package t31

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

// Конструктор
func NewPoint(x, y float64) *Point {
	return &Point{X: x,
		Y: y}
}

// Вычисление растояния(Евклидова метрика)
func (p *Point) Distance(p2 *Point) float64 {
	return math.Sqrt(math.Pow(p.X-p2.X, 2) + math.Pow(p.Y-p2.Y, 2))
}

func main() {

	a := NewPoint(2, 10)
	b := NewPoint(-1, 12)

	fmt.Println(a.Distance(b))

}
