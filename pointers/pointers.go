package main

import "fmt"

// Go supports pointers, allowing us to pass references to values and records within our program

// zeroval() has its argument passed by value, so it gets its own copy
func zeroval(ival int) {
	ival = 0
}

// zeroptr() has its argument as a pointer, which allows us to reference the value that the pointer
// points to -- i.e. it doesn't have its own copy of the value (but it does have its own local copy
// of the pointer variable)
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("Initial:", i)

	zeroval(i)
	fmt.Println("After zeroval:", i)

	zeroptr(&i)
	fmt.Println("After zeroptr", i)

	fmt.Println("Memory address of i:", &i)
}
