package main

import "fmt"

// Go makes it possible to recover form a panic, by using the built-in recover function

// A recover can stop a panic from aborting the program and let it continue with execution instead

// An example of where this can be useful:
// - a server wouldn't want to crash if one of the client connections exhibits a critical errors
// - instead, the server would want to close that connection and continue serving other clients
// - In fact, this is what Go's net/http does by default for HTTP servers

// This function panics
func mayPanic() {
	panic(
		"I'm panicking there appears to be a problem please help me dearly I am in great need of assistance please provide help at your soonest convenience",
	)
}

func main() {
	// recover must be called within a deferred function
	// when the enclosing function panics, the defer will activate and a receover call within it
	// will catch the panic

	// the return value of recover is the error raised in the call to panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered -- error:\n\t", r)
		}
	}()

	mayPanic()

	// this code will not run, beacause mayPanic panics
	// so the exeuction of main stops at the point of the panic, and resumes in the deferred closure
	fmt.Println("After myPanic()")
}
