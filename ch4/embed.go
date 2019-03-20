//Embedded member of struct
package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point  //embeded
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel
	w.X = 8 //advantage of embeded
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20

	w = Wheel{Circle{Point{8, 8}, 5}, 20}
	w = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20, // NOTE: trailing comma necessary here (and at Radius)
	}

	fmt.Printf("%#v\n", w) //%#v means use go grammer

	w.X = 42 //if Circle -> circle; it only can be access in package. But out of package, w.X can not be access.
	fmt.Printf("%#v\n", w)
}
