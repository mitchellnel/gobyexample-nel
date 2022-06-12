package main

import "fmt"

// Variadic functions can be called with any number of trailing arguments

// For example, fmt.Println() is a common variadic function

// Here is a function that will take an arbitrary number of ints as arguments
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	// We call variadic functions in the usual ways, passing as many parameters as we want
	sum(1, 2)
	sum(1, 2, 3)

	// If we already have multiple args in a slice, we can apply them to a variadic function using
	// 		function(slice...)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
