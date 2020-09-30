package exercise5

import (
	"fmt"
	"math"
)

// Exercise5 ... code for ninja level 6 exercise 5
func Exercise5() {
	fmt.Println("\nExercise 5:")

	sq := square{
		length: 4.5,
	}

	ci := circle{
		radius: 2,
	}

	re := rectangle{
		length: 4,
		width:  7,
	}

	fmt.Println("Square area:", info(sq))
	fmt.Println("Circle area:", info(ci))
	fmt.Println("Rectangle area:", info(re))
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	length float64
	width  float64
}

func (s square) area() float64 {
	return math.Pow(s.length, 2)
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r rectangle) area() float64 {
	return r.length * r.width
}

type shape interface {
	area() float64
}

func info(s shape) float64 {
	return s.area()
}
