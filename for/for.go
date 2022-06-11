package main

import "fmt"

func main() {
	// for loops are Go's only looping construct

	// Most basic type
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	// Classic for loop
	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}

	// We can use the 'break' keyword to break a loop
	for {
		fmt.Println("Oh no I'm stuck!")
		break
	}
	fmt.Println("I've been broken out of the infinite loop!")

	// We can use the 'continue' keyword to move onto the next iteration of the loop
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}

		fmt.Println(n)
	}

	// We’ll see some other for forms later when we look at range statements, channels, and other
	// data structures.
}
