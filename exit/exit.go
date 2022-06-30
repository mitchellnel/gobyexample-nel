package main

import (
	"fmt"
	"os"
)

// Use os.Exit to immediately exit with a given status

func main() {
	// defers will **not** be run when using os.Exit, so this fmt.Println will never be called
	defer fmt.Println("!")

	// exit with status 3
	os.Exit(3)

	// Note that unlike C, Go does not use an integer return value from main to indicate exit status
	// To exit with a nonzero status, you should use os.Exit
}
