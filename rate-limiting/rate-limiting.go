package main

import (
	"fmt"
	"time"
)

// Rate Limiting is an imporant mechanism for controlling resource utilisation and maintaining
// quality of service

// Go elegantly supports rate limiting with goroutines, channels and tickers

func main() {
	// first we'll look at basic rate limiting

	// suppose that we want to limit our handling of incoming requests
	// we'll serve these requests off a channel of the same name
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// this limiter channel will receive a value every 200ms
	// this is the regulator in our rate limiting scheme
	limiter := time.Tick(200 * time.Millisecond)

	// by blocking on a receive from the limiter channel before serving each request, we limit
	// ourselves to 1 request every 200ms
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	fmt.Println()

	// we may want to allow short bursts of request in our rate limiting scheme while preserving the
	// overall rate limit
	// we can accomplish this by buffering our limiter channel

	// this burstyLimited channel will allow bursts of up to 3 events
	burstyLimiter := make(chan time.Time, 3)

	// fill up the channel to represent allowed bursting
	for i := 1; i <= 3; i++ {
		burstyLimiter <- time.Now()
	}

	// every 200ms, we'll try to add a new value to burstyLimiter, up to its limit of 3
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// now simulate 5 more incoming requests
	// the first 3 of these will benefit from the burst capability of burstyLimiter
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

	// When we run our program:
	// - we see that the first batch of requests is handled once every ~200ms as desired
	// - we see that for the second batch requests, the first 3 are served immediately because of
	// the burstable rate limitng, then the remaining are served with ~200ms delays each
}
