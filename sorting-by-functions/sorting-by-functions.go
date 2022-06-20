package main

import (
	"fmt"
	"sort"
)

// Sometimes we'll want to sort a collection by something other than its natural order

// For example, suppose we wanted to sort strings by their length isntead of alphabetically

// Here's an example of custom sorts in Go

// In order to sort by a custom function in Go, we need a corresponding type
// Here, we create a byLength type that is just an alisa for the builtin []string type
type byLength []string

// We implement sort.Interface - Len, Less and Swap - on our type so we can use the sort package's
// generic Sort function

// Len and Swap will usually be similar across types
func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less will hold the actual custom sorting logic
// In our case, we want to sort in order of increasing string length, so we use len(s[i]) and
// len(s[j]) here
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	// with all of this in place, we can now implement our custom sort by converting the original
	// fruits slice to byLength, and then use sort.Sort on that typed slice
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)

	// By following this same pattern of creating a custom type, implementing the three Interface
	// methods on that type, and then calling sort.Sort on a collection of that custom type, we can
	// sort Go slices by arbitrary functions
}
