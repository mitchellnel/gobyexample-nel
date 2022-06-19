package main

import (
	"fmt"
	"time"
)

// We can use channels to synchronise execution across goroutines

// Here's an example of using a blocking receive to wait for a goroutine to finish

// When waiting for multiple goroutines to finish, you may prefer to use a **WaitGroup**

// This is the function that we'll run in the goroutine
// The done channel will be used to notify another goroutine that this function's work is done
func worker(done chan bool) {
	fmt.Print("Working ...")
	time.Sleep(time.Second)
	fmt.Println("... Done!")

	// send a value to the channel to notify that we're done
	done <- true
}

func main() {
	// start a worker goroutine, giving it the channel to notify on
	done := make(chan bool, 1)

	go worker(done)

	// block until we receive a notification from the worker on the channel
	<-done

	fmt.Println("All finished!")

	// If we removed the <-done line from the program, the program would exit before the worker even
	// started
}
