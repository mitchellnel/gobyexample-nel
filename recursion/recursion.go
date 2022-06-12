package main

import "fmt"

// Recursive factorial function
func fact(n int) int {
	// base case
	if n == 0 {
		return 1
	}

	// recursive step
	return n * fact(n-1)
}

func main() {
	fmt.Println("7! =", fact(7))

	// Closures can also be recursive, but this requires the closure to be declared with a typed var
	// explicitly before it is defined
	var fib func(n int) int

	fib = func(n int) int {
		// base case
		if n < 2 {
			return n
		}

		// recursive step
		// since fib was previously declared in main, Go knows which function to call with fib here
		return fib(n-1) + fib(n-2)
	}

	fmt.Println("fib(7):", fib(7))
}
