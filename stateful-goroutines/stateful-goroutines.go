package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// Another option to synchronise access to shared state across multiple goroutines is to use the
// built-in synchronisation features of goroutines and channels

// This channel-based approach aligns with Go's idea of sharing memory by communicating and having
// each piece of data owned by exactly 1 goroutine

// In this example, our state will be owned by a single goroutine
// This will guarantee that the data is never corrupted with concurrent access

// In order to read or write that state, other goroutines will send messages to the owning goroutine
// and receive corresponding replies

// These readOp and writeOp structs encapsulate those requests and a way for the owning goroutine to
// respond
type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	// as before, we'll count how many operations we perform
	var readOps uint64
	var writeOps uint64

	// the reads and writes channels will be used by other goroutines to issue read and write
	// requests, respectively
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// here is the goroutine that owns the state, which is a map as in the previous example, but the
	// map is now private to the stateful goroutine
	go func() {
		var state = make(map[int]int)

		// this goroutine repeatedly selects on the reads and writes channels, responding to
		// requests as they arrive
		// a response is executed by first performing the requested operation and then sending a
		// value on the response channel resp to indicate success (and the desired value in the case
		// of reads)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// this starts 100 goroutines to issue reads to the state-owning goroutine via the reads channel
	for r := 0; r < 100; r++ {
		go func() {
			for {
				// each read requires constructing a readOp
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}

				// sending the readOp over the reads channel
				reads <- read

				// receiving the result over the provided resp channel
				<-read.resp

				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// we start 10 writes as well, using a similar approach
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}

				writes <- write

				<-write.resp

				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// let the goroutines work for a second
	time.Sleep(time.Second)

	// finally, capture and report the op counts
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)

	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	// Running our program shows that the goroutine-based state management example completes
	// ~80k total operations

	// For this particular case, the goroutine-based approach was a bit more involved than the
	// mutex-based one It might be useful in certain cases though, for example where you have other
	// channels involved or when managin multiple such mutexed would be error-prone
}
