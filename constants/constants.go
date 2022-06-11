package main

import (
	"fmt"
	"math"
)

// The 'const' keyword declares a constant value
const C string = "Constant"

func main() {
	fmt.Println(C)

	// A const statement can appear anywhere that a var can
	const n = 500_000_000

	// Constant expressions perform arithmetic with arbitrary precision
	const d = 3e20 / n
	fmt.Println(d)

	// A numeric constant has no type until it is given one - such as by an explicit conversion
	fmt.Println(int64(d))

	// A numeric constant can be given a type by using it in a context that requires one, such as
	// variable assignment or a function call
	// For example, here math.Sin expects a float64, so d is given a type of float64 in this call
	fmt.Println(math.Sin(d))
}
