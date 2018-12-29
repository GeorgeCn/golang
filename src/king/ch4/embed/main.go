package main

import "fmt"

type Point struct {
	x, y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

var w Wheel
var w_ext Wheel
func main() {
	w = Wheel{Circle{Point{8, 8}, 5}, 20}
	fmt.Println(w)
	w_ext = Wheel{
		Circle:Circle{
			Point:Point{8, 8},
			Radius:5,
		},
		Spokes: 20,
	}
	fmt.Printf("%v#\n", w_ext)
	w.x = 42
	fmt.Println(w)
	w_ext.x = 42
	fmt.Println(w_ext)
}
