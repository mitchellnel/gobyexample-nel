package main

import "fmt"

// Go has built-in support for multiple return values

// This feature is often used in idiomatic Go, for example to retun both result and error values
// from a function

// The (int, int) in this function signature shows that the function returns two ints
func vals() (int, int) {
	return 3, 7
}

func main() {
	// Here, we use the 2 different return values from the call with multiple assignment in the one
	// line
	a, b := vals()
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	// If we only want a subset of the returned values, we use the blank identifier _
	_, c := vals()
	fmt.Println("c:", c)
}
