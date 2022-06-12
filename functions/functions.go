package main

import "fmt"

// Functions are central in Go

// Here's a function that takes in two ints, and returns their sum as an int
func plus(a int, b int) int {
	return a + b
}

// When you have multiple consecutive parameters of the same type,  you may omit the type name for
// the like-types parameters up to the final parameter that declares the type
func plusPlus(a, b, c int) int {
	return a + b + c
}

// Go requires explicit returns -- i.e. it won't automatically return the value of the last
// expression

func main() {
	// Call a function as we expect
	res := plus(1, 2)
	fmt.Println("1 + 2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 =", res)
}
