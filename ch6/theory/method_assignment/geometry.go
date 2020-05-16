package main

import (
	"fmt"
	"math"
)

// Point ...
type Point struct {
	X, Y float64
}

// Path ...
type Path []Point

func main() {
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}

	fmt.Println(perim.Distance())
}

// Distance функция
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance метод
// p - получатель метода приемника
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance ...
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		} // {1, 1}.Distance({5, 1})
	}
	return sum
}
