package main

import (
	"fmt"
	"math"
)

// Interfaces are named collections of method signatures

// Here's a bassic interface for geometric shapes
type geometry interface {
	area() float64
	perimeter() float64
}

// We will implement this interface on rect and circle types
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// To implement an interface on Go on some type, we just need to implement all of the methods in the interface

// Here, we implement geometry on rect
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return r.width*2 + r.height*2
}

// Here, we implement geometry on circle
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return math.Pi * 2 * c.radius
}

// If a variable has an interface type, then we can call methods that are in the named interface
// Here's a generic measure function taking advantage of this to work on any geometry type
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
