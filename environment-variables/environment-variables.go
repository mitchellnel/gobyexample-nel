package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// to set a key-value pair, use os.Setenv

	// to get a value for a key, use os.Getenv
	// this will return an empty string if the key isn't present in the environment

	os.Setenv("FOO", "1")

	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	// use os.Environ to list all key-value pairs in the environment
	// this returns a slice of string sin the form KEY=value
	fmt.Println()
	for _, e := range os.Environ() {
		// you can strings.SplitN them to get the key and value
		// here we print all keys
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[1])
	}
}
