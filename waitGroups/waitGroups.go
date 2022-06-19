package main

import (
	"fmt"
	"sync"
	"time"
)

// To wait for multiple goroutines to finish, we can use a wait group

// This is the worker function we'll run in every goroutine
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	// sleeping to simulate an expensive task
	time.Sleep(time.Second)

	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// this WaitGroup is used to wait for all the goroutines launched here to finish
	// note: if a WaitGroup is explicitly passed into functions, it should be done by _pointer_
	var wg sync.WaitGroup

	// launch several goroutines and increment the WaitGroup counter for each
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		// avoid reusing the same i value in each goroutine closure
		i := i
		// we have to do this because each iteration of the loop uses the same instance of the
		// variable i if we pass the worker goroutines the loop's instance of i, then each goroutine
		// will utilise whatever value of i the loop currently has

		// to avoid this, we have to bind the current value of i to each closure as it is launched
		// we do this by using i := i, which looks odd, but works fine in Go
		// it essentially creates a new instance of i for use in the closure

		// wrap the worker call in a closure that makes sure to tell the WaitGroup that this worker
		// is done this way, the worker itself does not have to be aware of the concurrency
		// primitives involved in its execution
		go func() {
			// the defer keywords defers the execution of a function until the surrounding function
			// returns
			defer wg.Done()
			worker(i)
		}()
	}

	// block until the WaitGroup goes back to 0; all the workers notified that they're done
	wg.Wait()

	// Note that this approach has no straightforward way to propagate errors from workers
	// Consider using the errgroup package to achieve this

	// The order of workers starting up and finishing is non-deterministic
}
