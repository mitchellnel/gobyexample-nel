package main

import "fmt"

// We previously saw how for and range provide iteration over basic data structures

// We can also use this syntax to iterate over values received from a channel

func main() {
	// we'll iterate over two values in the queue channel
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// this range iterates over each element as it's received from queue
	for elem := range queue {
		fmt.Println(elem)
	}
	// because we closed the channel above, the iteration terminates after receiving the 2 elements

	// This example also showed that it's possible to close a non-empty chanel, but still have the
	// remaining values be received
}
