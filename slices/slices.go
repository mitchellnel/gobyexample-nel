package main

import "fmt"

// slices are a key data type in Go, giving a more powerful interfact to sequences than arrays do

func main() {
	// Unlike arrays, slices are typed only by the elements they contain, not the number of elements

	// To create an empty slice with non-zero length, we use the builtin make
	// Here, we make a slice of length 4 of strings, all zero-initialised
	s := make([]string, 4)
	fmt.Println("emp:", s)

	// We can set and get just like we do with arrays
	s[0] = "Lewis Hamilton"
	s[1] = "George Russell"
	s[2] = "Charles Leclerc"
	s[3] = "Carlos Sainz"
	fmt.Println("drivers:", s)
	fmt.Println("Monagasque:", s[2])

	// We use len to get the length
	fmt.Println("len:", len(s))

	// In addition to these basic operations, slices support several more methods that make them
	// richer than arrays

	// One is the built-in append(), which returns a slice containing one or more new values
	// Note that we need to accept a return value from append as we may get a new slice value
	s = append(s, "Max Verstappen", "Sergio Perez")
	fmt.Println("drivers apd:", s)

	// Slices can also be copy'd
	c := make([]string, len(s))
	copy(c, s) // note the arguments are dst, src
	fmt.Println("cpy:", c)

	// Slices support a "slice" operator with the syntax
	// 		slice[low:high]
	// where we get slice[low], slice[low+1], ..., slice[high-1]
	l := s[2:5]
	fmt.Println("sl1:", l)

	// This slices up to, but excluding, the high argument
	l = s[:5]
	fmt.Println("sl2:", l)

	// And this slices up from, and including, the low argument
	l = s[2:]
	fmt.Println("sl3:", l)

	// We can declare and initialise a variable for slice in a single line as well
	t := []string{"Alexander Albon", "Lando Norris", "George Russell"}
	fmt.Println("2019 Rookies:", t)

	// Slices can also be composed into multi-dimensional data structures
	// The length of the inner slices can vary, unlike with multi-dimensional arrays
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)

		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	// Again, fmt.Println() pretty prints slices like arrays
}
