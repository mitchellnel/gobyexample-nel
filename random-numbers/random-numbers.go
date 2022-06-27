package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Go's math/rand package provides pseudorandom number generation

func main() {
	// the default number generator is deterministic, so it'll produce the same sequence of numbers
	// each time by default
	// to produce varying sequences, give it a seed that changes

	// note that this is not safe to use for random numbers you intend to be secret
	// use crypto/rand for those

	// for the functions to produce non-deterministic numbers we have to set a seed
	rand.Seed(time.Now().UnixNano())

	// rand.Intn(arg) returns a random int n, 0 <= n < arg
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	// rand.Float64(seed) returns a float64 f, 0.0 <= f < 1.0
	fmt.Println(rand.Float64())

	// this can be used to generate random floats in other ranges
	// for eaxmple 5.0 <= f' < 10.0
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// instead of seed setting we can also do
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// then call the functions our our rand object
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	// if you seed a source with the same number, it produces the same sequence of random numbers
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()

	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
	fmt.Println()
}
