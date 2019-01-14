package main

import "fmt"

// struct {
// 	var vartype;
// }

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

// Create a new struct for contains {X,Y}
type Point struct {
	X, Y int
}

// create a wheel
func main() {

	// version with out particular struch
	// var w Wheel
	// w.X = 8
	// w.Y = 8

	// w.Radius = 5
	// w.Spokes = 20

	// verbose access fields
	// var w newWheel
	// w.Circle.Center.X = 8
	// w.Circle.Center.Y = 8
	// w.Circle.Radius = 5
	// w.Spokes = 20

	// thanks to anonymous fields
	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20

	fmt.Printf("%#v\n", w)

}
