package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Path []Point

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func main() {
	p := Point{X: 1, Y: 2}
	q := Point{X: 4, Y: 6}
	fmt.Println(p.Distance(q))

	path := Path{
		{1, 1}, {5, 1}, {5, 4}, {1, 1},
	}
	fmt.Println(path.Distance())
}
