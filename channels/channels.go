package main

import "fmt"

// Channels are the pipes that connect concurrent goroutines

// You can send values into channels from one goroutine, and receive those values into another
// goroutine

func main() {
	// Create a new channel with
	//   make(chan val-type)
	// Channels are types by the value that they convey
	messages := make(chan string)

	// Send a value into a channel using the
	//   channel <- value
	// syntax

	// Here, we send "ping" to the messages channel from a new goroutine
	go func() { messages <- "ping" }()

	// The
	//   variable <-channel
	// syntax receives a value from the channel

	// Here, we'll receive the "ping" message that we sent above and print it out
	msg := <-messages
	fmt.Println(msg)

	// When we run the progrm, the "ping" message is successfully passed from one goroutine to
	// another via our channel

	// By default, sends and received block until both the sender and the receiver are ready
	// This property allowed us to wait at the end of our program for the "ping" message without
	// having to use any other synchronisation
}
