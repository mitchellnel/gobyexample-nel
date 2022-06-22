package main

import (
	"fmt"
	"os"
)

// Go offers excellent support for string formatting in the printff tradition

// Here are some common examples

var printf = fmt.Printf

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}

	// Go offers several printing "verbs" designed to format general Go values

	// This prints an instance of our point struct
	printf("struct1: %v\n", p)

	// If the value is a struct, the %+v variant will include the struct's field names
	printf("struct2: %+v\n", p)

	// The %#v variant prints a Go syntax representation of the value, i.e. the source code snippet
	// that would produce that value
	printf("struct3: %#v\n", p)

	// To print the type of a value, use %T
	printf("type: %T\n", p)

	// Formatting booleans is straightforward
	printf("bool: %t\n", true)

	// There are many options for formatting integers

	// Use %d for standard, base-10 formatting
	printf("int: %d\n", 123)

	// %b prints a binary representation
	printf("binary: %b\n", 10)

	// %x provides hex encoding
	printf("hex: %x\n", 44)

	// %c prints the character correspodning to the given integer
	printf("char: %c\n", 33)

	// there are also several formatting options for floats

	// %f provides basic decimal formatting
	printf("float1: %f\n", 78.9)

	// %e and %E format the float in slightly different version of scientific notation
	printf("float2: %e\n", 123400000.0)
	printf("float3: %E\n", 123400000.0)

	// %s provides basic string printing
	printf("str1: %s\n", "iamastring")

	// %q alloes us to double-quote strings as in Go source
	printf("str2: %q\n", "\"iamalsoastring\"")

	// %x will also render strings in base-16, with two output characters per byte of input
	printf("str3: %x\n", "anotherstring")

	// %p will print the representation of a pointer
	printf("pointer: %p\n", &p)

	// when formatting numbers, you will often want to control the width and precision of the
	// resulting figure

	// to specify the width of an integer, use a number after the % in the verb
	// by default the result will be right-justified and padded with spaces
	printf("width1: |%6d|%6d|\n", 12, 345)

	// you can also specify the width of printed floats, though usually you'll also want to restrict
	// the decimal precision at the same time with width precision syntax
	printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	// to left justify, use the - flag
	printf("width2: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// you may also want to control width when formatting strings, especially to ensure that they
	// align in table-like output

	// for basic right-justified width
	printf("width4: |%6s|%6s|\n", "foo", "b")

	// to left-justify, use the - flag as with numbers
	printf("width4: |%-6s|%-6s|\n", "foo", "b")

	// Sprintf formats and returns a string without printing it anywhere
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	// we can format and print to io.Writers other than os.Stdout using Fprintf
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
