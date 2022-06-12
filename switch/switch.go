package main

import (
	"fmt"
	"time"
)

func main() {
	// Basic switch statements
	i := 2
	fmt.Print("Write ", i, " as: ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// Note the absence of break statements

	// We can use commas to separate multiple expressions in the same case statement
	// We also have the default case statement
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday...'")
	}

	// switch without a given expression is an alternate way to express if/else logic
	// We also see that the case expressions can be non-constants
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// In Go, we can create a "type switch"
	// A type switch comapares types instead of values
	// We can use this to discover the type of an interface value
	// In this example, the variable t will have the type corresponding to its clause
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am an int")
		default:
			fmt.Printf("I don't know what type I am ... (maybe a %T)\n", t)
		}
	}

	whatAmI(true)
	whatAmI(69)
	whatAmI("Hello, World!")
	whatAmI(3.14159265)

	const n = 44
	whatAmI(n)
}
