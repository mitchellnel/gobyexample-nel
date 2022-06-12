package main

import "fmt"

// range iterates over elements in a variety of data structures

func main() {
	// Here, we use range to sum the numbers in a slice -- arrays work like this too
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("nums:", nums)
	fmt.Println("sum:", sum)

	// range on arrays and slices provides both the index and valuef or each entry
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index of 3:", i)
		}
	}

	// range on map iterates over key-value pairs
	kvs := map[string]int{
		"Lewis Hamilton": 44,
		"George Russell": 63,
		"Max Verstappen": 1,
		"Sergio Perez":   11,
	}
	for k, v := range kvs {
		fmt.Printf("%s -> %d\n", k, v)
	}

	// range can also iterate over just the keys of a map
	for k := range kvs {
		fmt.Println("Key:", k)
	}

	// range on strings iterates over Unicode code points
	// The first value is the starting byte index of the rune and the second the rune itself
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
