package main

import (
	"fmt"
	"unicode/utf8"
)

// A Go string is a read-only slice of bytes
// Go treates strings as containers of text encoded in UTF-8

// In other languages, strings are composed of characters
// In Go, the concept of a character is called a rune -- an integer that represents a Unicode code points
// https://go.dev/blog/strings

func main() {
	// s is a string assigned a literal value representing the word "hello" in Thai
	const s = "สวัสดี"
	// Recall Go strings are UTF-8 encoded text

	// Since strings are equivalent to []bytes, the below line produces the length of the raw bytes stored within
	fmt.Println("Len:", len(s))

	// Indexing into a string produces the raw byte values at each index
	// The below loop generates the hex values of all the bytes that constitute the code points in s
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// To count how many runes are in a astring, we can use the utf8 package
	fmt.Println("Rune count:", utf8.RuneCountInString(s))
	// The runtime of RuneCountInString depends on the size of the string as it decodes each UTF-8 rune sequentially
	// Some Thai characters are represented by multiple UTF-8 code points, so the result of this count may be surprising

	// A range loop handles strings specially, and decodes each rune along with its offset in the string
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// We can achieve the same iteration by using the utf8.DecodeRuneInString function explicitly
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		// We can pass rune values to functions
		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	// Values enclosed in single quotes are rune literals
	// We can compare a rune value to a rune literal directly
	if r == 't' {
		fmt.Println("Found tee")
	} else if r == 'ส' {
		fmt.Println("Found so sua")
	}
}
