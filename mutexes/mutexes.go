package main

import (
	"fmt"
	"sync"
)

// In the previous example, we saw how to manage simple counter state using atomic operations

// For more compelx state, we can use a mutex lock to safely access data across multiple goroutines

// Container holds a map of counters
// Since we want to update it concurrently from multiple goroutines, we add a Mutex to synchronise
// access
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

// Note that mutexes must not be copied, so if this struct is passed around, it should be done by
// pointer

// Increment function for the counter with a specific name
func (c *Container) inc(name string) {
	// lock the mutex before accessing counters
	c.mu.Lock()

	// unlock it at the end of the function using a defer statement
	defer c.mu.Unlock()

	c.counters[name]++
}

func main() {
	// note that the zero value of a mutex is useable as-is, so no mutex initialisation is required
	// here
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	// this function increments a named counter in a loop
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	// run several goroutines concurrently
	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)
	// note that they all access the same container c, and two of them access the same counter

	// wait for the goroutines to finish
	wg.Wait()

	fmt.Println(c.counters)

	// Running the program shows that the counters updates as expected
}
