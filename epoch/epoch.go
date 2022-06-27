package main

import (
	"fmt"
	"time"
)

// A common requirement in programs is getting the number of seconds, milliseconds or nanoseconds
// since the Unix epoch

func main() {
	// use time.Now with Unix, Unixmilli or UnixNano to get elapsed time since the Unix epoch in
	// seconds, milliseconds or nanoseconds, respectively
	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	// we can also convert integer seconds or nanoseconds since the epoch into the corresponding
	// time
	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}
