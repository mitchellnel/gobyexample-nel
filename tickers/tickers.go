package main

import (
	"fmt"
	"time"
)

// Timers are for when we want to do something once in the future

// Tickers are for when we want to do something repeatedly at regular intervals

func main() {
	// here's an example of a ticker that ticks periodically until we stop it

	// tickers use a similar mechanis to timers: a channel that is sent values
	// here we'll use the select built-in on the channel to await the values as they arrive every
	// 500ms
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// tickers can be stopped like timers
	// one a ticker is stopped it won't receive any more values on its channel

	// we'll stop our ticker after 1600ms
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()

	done <- true

	fmt.Println("Ticker stopped")

	// When we run this program, the ticker should tick 3 times before we stop it
}
