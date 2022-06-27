package main

import (
	"fmt"
	"strconv"
)

// Parsing numbers from strings is a basic but common task in many programs

// The built-in package strnconv provides the number parsing

func main() {
	// with ParseFloat, the 64 indicates how mayn bits of precision to use
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// for ParseInt, the 0 means infer the base form the string
	// 64 requires that the result fit in 64 bits
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// ParseInt also recognises hex formatted numbers
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// a ParseUint is also available
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// Atoi is a convenience function for basic base-10 int parsing
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// parse functions return an error on bad input
	_, err := strconv.Atoi("wat")
	fmt.Println(err)
}
