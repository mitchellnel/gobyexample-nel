package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// The primary mechanism for managing state in Go is communication over channels
// (for example in worker pools)

// There are a few other options for managing state though

// Here, we'll look at using the sync/atomic package for atomic counters accessed by multiple
// goroutines

func main() {
	// we'll use a uint to represent our counter
	var ops uint64

	// a WaitGroup will help us wait for all goroutines to finish their work
	var wg sync.WaitGroup

	// we'll start 50 goroutines that each increment the counter exactly 1000 times
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				// to atomically increment the counter, we use AddUint64, giving it the memory
				// address of our ops counter with the & syntax
				atomic.AddUint64(&ops, 1)
			}

			wg.Done()
		}()
	}

	// wait until all goroutines are done
	wg.Wait()

	// it's safe top access ops now because we know no other goroutine is writing to it
	fmt.Println("ops:", ops)

	// Reading atomics safely while they are being updated is also possible, using functions like
	// atomic.LoadUint64

	// When we run, we expect to get exactly 50,000 operations

	// Had we used the non-atomic ops++ to increment the counter, we would likely get a different
	// number, changing between runs, because the goroutines would interfere with each other

	// In fact, if you use ops++ and go run -race we will observe the existence of data race
	// failures
}
