package main

import (
	"bufio"
	"fmt"
	"os"
)

// Writing files in Go follows similar patterns to the ones we saw earlier for reading

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// here's how to dump a string, or just bytes, into a file
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// for more granual writes, open a file for writing
	f, err := os.Create("/tmp/dat2")
	check(err)

	defer f.Close() // it's idomatic to defer a Close immediately after opening a file

	// you can Write byte slices as you'd expect
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// a WriteString is also available
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// issue a Sync to flush write to stable storage
	err = f.Sync()
	check(err)

	// bufio provides buffered writers in addition to the buffered readers that we saw earlier
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered write\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	// use Flush to ensure all buffered operations have been applied to the underlying writer
	w.Flush()

	// Run the code, and then check the contents of the written files
}
