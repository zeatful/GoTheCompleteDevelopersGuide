package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func (t triangle) getArea() float64 {
	return (.5 * t.base * t.height)
}

func (s square) getArea() float64 {
	return (s.sideLength * s.sideLength)
}

func main() {
	s := square{}
	t := triangle{}

	s.sideLength = 5

	t.base = 5
	t.height = 2

	fmt.Println("Area of Square of ", s.sideLength, "x", s.sideLength, "(L*W) = ", s.getArea())
	fmt.Println("Area of Triangle of ", t.base, "x", t.height, "(.5*base*height) = ", t.getArea())
}
