package main

import (
	"os"
	"os/exec"
	"syscall"
)

// Sometimes, we just want to completely replace the current Go process with another, perhaps
// non-Go, one

// To do this, we use the Go's implementation of exec

func main() {
	// for our example, we'll exec ls

	// Go requires an absolute path to the binary that we want to execute, so we'll use
	// exec.LookPath to find it (probably /bin/ls)
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// exec requires arguments in slice form (as opposed to one big string)
	// we'll give ls a few common arguments
	// note that the first argument should be the program name
	args := []string{"ls", "-a", "-l", "-h"}

	// exec also needs a set of environment variables to use
	// here we just provide our current environment
	env := os.Environ()

	// here's the actual syscall.Exec call
	// if this call is successful, the execution of our process will end here and be replace by the
	// /bin/ls -a -l -h process
	// if there is an error, we'll get a return value
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}

	// Note that Go does not offer an equivalent for the fork function
	// Starting goroutines, spawning process and execing processes covers most of the use cases for
	// forks
}
