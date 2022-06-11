package main

import "fmt"

func main() {
	// Go will infer the type of initialized variables, unless specified otherwise

	var a string = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d bool = true
	fmt.Println(d)

	// Variables declared without a corresponding initialization are zero-valued. For example, the
	// zero value for an int is 0
	var e int
	fmt.Println(e)

	var f float64 = 3.14159265
	fmt.Println(f)

	// The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string =
	// "apple" in this case
	g := "apple"
	fmt.Println(g)
}
