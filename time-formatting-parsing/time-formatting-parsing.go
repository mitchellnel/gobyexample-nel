package main

import (
	"fmt"
	"time"
)

// Go supports time formatting and parsing via pattern-based layouts

var print = fmt.Println

func main() {
	// here's a basic example of formatting a time according to RFC3339, using the corresponding
	// layout constant
	t := time.Now()
	print(t.Format(time.RFC3339))

	// time parsing uses the same layout values as Format
	t1, _ := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	print(t1)

	// Format and Parse use example-based layouts
	// usually, we'll use a constant from time for these layouts, but we can also supply custom
	// layouts

	// layouts must use the reference time
	// 		Mon Jan 2 15:04:05 MST 2006
	// to show the pattern with which to format/parse a given time/string
	// the example time must be exactly as shown: the year 2006, 15 for the hour, Monday for the day
	// of the week etc.
	print(t.Format("3:04PM"))
	print(t.Format("Mon Jan _2 15:04:05 2006"))
	print(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 41 PM")
	print(t2)

	// for purely numeric respresentations you can also use the standard string formatting with the
	// extraced components of the time value
	fmt.Printf(
		"%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
	)

	// parse will return an error on malformed input explaining the parsing problem
	ansic := "Mon Jan _2 15:04:05 2006"
	_, err := time.Parse(ansic, "8:41PM")
	print(err)
}
