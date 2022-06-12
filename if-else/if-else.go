package main

import "fmt"

func main() {
	// Basic example
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// Normal if statement
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// A statement can precede the conditionals in an if-statement
	// If any variables are declared in this statement, then that variable is available to all
	// branches of the if-block
	if num := 9; num < 0 {
		fmt.Println(num, "is negative") // space automatically put between arguments to fmt.Println
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has more than 1 digit")
	}

	// Note that brackets are not needed around the predicates

	// There is no ternary operator in Go -- if-statements must also be used for basic conditionals
}
