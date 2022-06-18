package main

import "fmt"

// Go supports methods defined on struct types

type rect struct {
	width, height int
}

// This area method has a receiver type of *rect
func (r *rect) area() int {
	return r.width * r.height
}

// Methods can be defined for either value or pointer receiver types
func (r rect) perimeter() int {
	return r.width*2 + r.height*2
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area:", r.area())
	fmt.Println("perimeter:", r.perimeter())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perimeter:", rp.perimeter())

	// Go automatically handles conversion betweeen values and pointers for method calls
	// We may want to use a pointer receiver type to avoid copying on method calls or to wallow the
	// method to mutate the receiving struct

	// But we do not have to change the way we call the method
}
