package main

import "fmt"

// In Go, an array is a numbered sequence of elements of a specific length

func main() {
	// Here, we create an array that will hold 5 ints
	var arr [5]int
	fmt.Println("arr:", arr)
	// The type of the elements and the length of the array are both part of its type
	// By default, an array is zero-values -- e.g. 0 for ints and "" for strings

	// We use the typical syntax to set and get specific elements
	arr[4] = 33
	fmt.Println("Set:", arr)
	fmt.Println("Get:", arr[4])

	// Use len() to return the length of the array
	fmt.Println("arr is of length", len(arr))

	// Use this syntax to declare and initialise an array in one line
	arr2 := [6]int{3, 33, 4, 44, 5, 55}
	fmt.Println("arr2:", arr2)

	// Array types are one-dimensional, but as in C, we can compose types to create  multi-dimensional arrays
	var twoD [2][2]int
	for i := 1; i <= 2; i++ {
		for j := 10; j >= 9; j-- {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD:", twoD)

	// Note that arrays appear in a nice prettified format when printed with fmt.Println()

	// We see slices much more than typical arrays in Go
}
