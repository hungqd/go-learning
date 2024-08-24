package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel
	w.X = 10
	w.Y = 10
	w.Radius = 5
	w.Spokes = 20
	fmt.Println(w)

	// no shorthand syntax
	w = Wheel{
		Circle: Circle{
			Point:  Point{X: 3, Y: 3},
			Radius: 6,
		},
		Spokes: 15,
	}
	fmt.Println(w)
}
