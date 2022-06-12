package main

import "fmt"

// Maps are Go's built-in associative data type

func main() {
	// To create an empty map, we use the builtin make
	// 		make(map[key-type]value-type)
	m := make(map[string]int)

	// We set and get like we do in Python
	m["Lewis Hamilton"] = 44
	m["George Russell"] = 63

	fmt.Println("Driver Numbers:", m)

	grNum := m["George Russell"]
	fmt.Println("George Russell drives with the number", grNum)

	// len() returns the number of key-value pairs in the map
	fmt.Println("len:", len(m))

	// The builtin delete() removes key-value pairs from the map
	delete(m, "George Russell")
	fmt.Println("Driver Numbers:", m)

	// The optional second return value when getting a value from a map indicates if the key was
	// present in the map
	// This can be used to disambiguate between missing keys, and keys with zero values like 0 or ""
	_, present := m["George Russell"]
	fmt.Println("present:", present)
	// Here, we didn't need the value itself, so we ignored it with the blank identifier "_"

	// We can also declare and initialise a map in one line
	n := map[string]int{"Charles Leclerc": 16, "Carlos Sainz": 55}
	fmt.Println("Ferrari Driver Numbers:", n)

	// As with before, maps are (somewhat) pretty printed when using fmt.Println()
}
