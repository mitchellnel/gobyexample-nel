package main

import "embed"

// `//go:embed` is a compiler directive that alows programs to include arbitrary files and folders in
// the Go binary at build time

// you can read more here:
//   https://pkg.go.dev/embed

// embed directives accept paths relative to the directory containing the Go source file

// this directive embeds the contents of the file into the string variable immediately following it
//go:embed folder/single_file.txt
var fileString string

// or embed the contents of a file into a []byte
//go:embed folder/single_file.txt
var fileByte []byte

// we can also embed multiple files or even folders with wildcards

// this uses a variable of the embed.FS type, which implements a simple virtual file system
//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {
	// print out the conttents of single_file.txt
	print(fileString)
	print(string(fileByte))

	// retrieve some files from the embedded folder
	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}
