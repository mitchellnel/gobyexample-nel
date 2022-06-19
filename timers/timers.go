package main

import (
	"fmt"
	"time"
)

// We often want to execute Go code at some point in the future, or repeatedly at some interval

// Go's built-in timer and ticker features make both of these tasks easy

// We'll look first at timers, and then at tickers

func main() {
	// timers represent a single event in the future
	// you tell the timer how long you want to wait, and it provides a chanel that will be notified
	// at that time

	// this timer will wait 2 seconds
	timer1 := time.NewTimer(2 * time.Second)

	// the <-timer1.C blocks on the timer's channel C until it sends a value indicating that the
	// timer has fired
	<-timer1.C
	fmt.Println("timer1 fired")

	// if we only wanted to wait, we could've just used timer.Sleep

	// but one reason a timer may be useful is that you can cancel the timer before it fires
	// here is an example of that
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer2 stopped")
	}

	// give timer2 enough time to fire, if it ever was going to, to show that it is in fact stopped
	time.Sleep(2 * time.Second)

	// The first timer will fire after ~2s after we start the program, but the second should be
	// stopped before it has a chance to fire
}
