package main

import (
	"fmt"
	"time"
)

// Go's select lets you wait on multiple channel operations

// Combining goroutines and channels with select is a powerful feature of Go

func main() {
	// for our example, we'll select across two channels
	c1 := make(chan string)
	c2 := make(chan string)

	// each channel will receive a value after some amount of time, to simulate e.g. blocking RPC
	// operations executing in concurrent goroutines
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		// we'll use select to await both of these values simultaneously, printing each one as it
		// arrives
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	// We then receive the values "one" and "two" as expected

	// If we
	//   $ time go run .
	// then we note that the total execution time is only ~2 seconds, since both the 1 second and
	// the 2 second sleeps execute concurrently
}
