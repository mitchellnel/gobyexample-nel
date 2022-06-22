package main

import (
	"fmt"
	"os"
)

// Defer is used to ensure that a function call is performed later in a program's execution, usually
// for the purposes of cleanup

// defer is often used where e.g. ensure and finally would be used in other languages

func main() {
	// suppose that we wanted to create a file, write to it, and then close it when we're done

	// here's how we could do that with defer
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)

	// immediate after getting a file object with createFile, we defer the closing of that file with
	// closeFile
	// this will be executed at the end of the enclosing function (main), after writeFile() has
	// finished
}

func createFile(filename string) *os.File {
	fmt.Println("Creating", filename)
	f, err := os.Create(filename)

	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("Writing to file", f.Name(), "...")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("Closing", f.Name(), "...")
	// it's important to check for errors when closing a file, even in a deferred function
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
