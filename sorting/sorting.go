package main

import (
	"fmt"
	"sort"
)

// Go's sort package implements sorting for builtins and user-defined types

func main() {
	// sort methods are specific to the builtin type

	// here's an example for strings
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)
	// note that sorting is in-place -- so it changes the given slice and doesn't return a new one

	// sorting ints
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:", ints)

	// we can also use sort to check if a slice is already in sorted order
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)
}
