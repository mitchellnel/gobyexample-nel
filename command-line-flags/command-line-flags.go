package main

import (
	"flag"
	"fmt"
)

// Go provides a flag package supporting basic command-line flag parsing

// We'll use this package to implement our example command-line program

func main() {
	// basic flag declaration are available for string, integer and Boolean options

	// here, we declare a string flag word with a default value "foo" and a short description
	// this flag.String function returns a string pointer
	wordPtr := flag.String("word", "foo", "a string")

	// this declares numb and fork flags using a similar apporach to the word flag
	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	// it's also possible to declare and option that uses an existing var declared elsewhere in the
	// program
	// to do this, we pass in a pointer to the flag declaration function
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// once all flags are declared, call flag.Parse() to execute command-line parsing
	flag.Parse()

	// here we'll just dump out the parsed options and and any trailing positional arguments
	// note that we need to dereference the pointers to get the actual option values
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())

	// To experiment with command-line arguments, it's best to build a binary with go build first
	// >> go build command-line-flags.go

	// Try out the build program by first giving it values for all flags
	// >> ./command-line-flags -word=opt -numb=7 -fork -svar=flag

	// Note that if you omit flags they automatically take their default values
	// >> ./command-line-flags -word=opt

	// Trailing positional arguments can be provided after any flags
	// >> ./command-line-flags -word=opt a1 a2 a3

	// Note that the flag package requries all flags to appear before positional arguments --
	// otherwise the flags will be interpreted as positional arguments
	// >> ./command-line-flags -word=opt a1 a2 a3 -numb=7

	// Use -h or --help flags to get automatically generated help text for the command-line program
	// >> ./command-line-flags -h

	// If you provide a flag that wasn't specified to the flag package, the program will print an
	// error message and show the help text again
	// >> ./command-line-flags -wat
}
