package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base float64
	height float64
}

type square struct {
	sideLength float64
}

func main() {
	t := triangle{base: 3, height: 4}
	s := square{sideLength: 8}

	printArea(t)
	printArea(s)
}

func (t triangle) getArea() float64 {
	return t.base * t.height / 2
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}