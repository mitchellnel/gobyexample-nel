package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args provides access to raw command-line arguments
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	// note that the first value in the os.Args slice is the path to the program, and os.Args[1:]
	// holds the arguments to the program

	// we can get individual args with normal indexing
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)

	// To experiment with command-line arguments, it's best to build a binary with go build first
	// >> go build command-line-arguments.go
	// ./command-line-arguments a b c d
}
