package main

// The standard library's strings package provides many useful string-related functions

// Here are some examples to give us a sense of the package

import (
	"fmt"
	s "strings"
)

// Aliasing fmt.Println
var print = fmt.Println

func main() {
	// below is a sample of the functions available in strings
	print("Contains:  ", s.Contains("test", "es"))
	print("Count:     ", s.Count("test", "t"))
	print("HasPrefix: ", s.HasPrefix("test", "te"))
	print("HasSuffix: ", s.HasSuffix("test", "st"))
	print("Index:     ", s.Index("test", "e"))
	print("Join:      ", s.Join([]string{"a", "b"}, "-"))
	print("Repeat:    ", s.Repeat("a", 5))
	print("Replace:   ", s.Replace("foo", "o", "0", -1))
	print("Replace:   ", s.Replace("foo", "o", "0", 1))
	print("Split:     ", s.Split("a-b-c-d-e", "-"))
	print("ToLower:   ", s.ToLower("TEST"))
	print("ToUpper:   ", s.ToUpper("test"))

	// Since these are functions from the package, and not methods on the string object itself, we
	// need to pass the string in question as the first argument to the function

	// https://pkg.go.dev/strings
}
