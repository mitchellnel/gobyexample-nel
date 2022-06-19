package main

import (
	"fmt"
	"time"
)

// A goroutine is a lightweight thread of execution

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// Suppose that we have a function call f(s)
	// Here's how we'd call it in the usual way, running it synchronously
	f("direct")

	// To invoke this function in a goroutine, we use:
	//   go f(s)
	// This new goroutine will execute concurrently with the calling one
	go f("goroutine")

	// We can also start a goroutine for an anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Our two function calls are running asynchronously in separate goroutines now
	// Wait for them to finish (use a WaitGroup for a more robust approach)
	time.Sleep(time.Second)
	fmt.Println("done")

	// When we run the program, we will deterministically see the output of the blocking call first,
	// and then the output of the two goroutines

	// the goroutines' output may be interleaved, because goroutines are being run concurrently by
	// the Go runtime
}
