package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Sometimes, we'd like our Go programs to intelligently handle Unix signals

// For example, we might want a server to gracefully shutdown when it receives a SIGTERM, or a
// command-line tool to stop processing input if it receives a SIGINT

// We can handle signals in Go with channels

func main() {
	// Go signal notification works by sending os.Signal values on a channel
	// we'll create a channel to receive these notifications
	// note that this channel should be buffered
	sigs := make(chan os.Signal, 1)

	// signal.Notify registers the given channel to receive notifications of the specified signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// we could receive from sigs here in the main function, but let's see how this could also be
	// done in a separate goroutine to demonstrate a more realistic scenario of graceful shutdown
	done := make(chan bool, 1)

	// this goroutine executes a blocking receive for signals
	// when it gets one, it'll print it out and then notify the program so that it can finish
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	// the program will wait here until it gets the expected signal
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
