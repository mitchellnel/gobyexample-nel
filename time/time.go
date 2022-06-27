package main

import (
	"fmt"
	"time"
)

// Go offers extensive support for times and durations
var print = fmt.Println

func main() {
	// getting the current time
	now := time.Now()
	print(now)

	// we can build a time struct by providing the year, month, day etc.
	// times are always associated with a Location, i.e. timezone
	then := time.Date(2009, 3, 12, 21, 21, 8, 694206, time.UTC)
	print(then)

	// we can then extract various components of the time value as expected
	print(then.Year())
	print(then.Month())
	print(then.Day())
	print(then.Hour())
	print(then.Minute())
	print(then.Second())
	print(then.Nanosecond())
	print(then.Location())

	// the Monday-Sunday Weekday is also available
	print(then.Weekday())

	// these methods compare two times, testing if the first occurs before, after, or at the same
	// time as the second, respectively
	print(then.Before(now))
	print(then.After(now))
	print(then.Equal(now))

	// The Sub methods returns a Duration representing the interval between two times
	diff := now.Sub(then)
	print(diff)

	// we can compute the length of the duration in various units
	print(diff.Hours())
	print(diff.Minutes())
	print(diff.Seconds())
	print(diff.Nanoseconds())

	// we can use Add to advance a time by a given duration, or with a - to move backwards by a
	// duration
	print(then.Add(diff))
	print(then.Add(-diff))
}
