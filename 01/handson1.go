package main

import (
	"fmt"
)

type shape interface {
	area()
}

type circle struct {
	radius int
}

func (c circle) area() {
	length := 3.14 * float32(c.radius) * 2.0
	area := length * length
	fmt.Println("The area of this circle is", area)
}

type square struct {
	side int
}

func (s square) area() {
	area := s.side * s.side
	fmt.Println("The area of this square is", area)
}

func printArea(s shape) {
	s.area()
}

func main() {
	c := circle{2}
	s := square{5}

	printArea(c)
	printArea(s)
}
