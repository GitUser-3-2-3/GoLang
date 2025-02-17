package main

import "fmt"

type shape interface {
	perimeter() float64
	area() float64
}

type rect struct {
	width  float64
	height float64
}

func (r rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (r rect) area() float64 {
	return r.width * r.height
}

type circle struct {
	radius float64
}

func (c circle) perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func getPerimeter(s shape) float64 {
	return s.perimeter()
}

func getArea(s shape) float64 {
	return s.area()
}

func main() {
	r := rect{width: 10, height: 20}
	c := circle{radius: 10}
	fmt.Print("Perimeter of rectangle: ", getPerimeter(r), "\n")
	fmt.Print("Perimeter of circle: ", getPerimeter(c), "\n")
}
