package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

// The filepath package provides functions to parse and construct file paths in a way that is
// portable between operating systems

// dir/file on Linux vs dir\file on Windows, for example

func main() {
	// Join should be used to construct paths in a portable way
	// it takes any number of arguments and construct a hierarchical path from them
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// we should always use Join instead of concatenating manually
	// in addition to providing portability, Join will also normalise paths by removing superfluous
	// separators and directory changes
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// Dir and Base can be used to split a path to the directory and the file
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	// alternatively, Split will return both in the same call
	dir, base := filepath.Split(p)
	fmt.Println("Split(p):", dir, base)

	// we can also check whether a path is absolute
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	// some file names have extensions following a dot
	// we can split the extension out of such names with Ext
	filename := "config.json"

	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// to find the file's name with the extension removes, use strings.TrimSuffix
	fmt.Println(strings.TrimSuffix(filename, ext))

	// Rel finds a relative path between a base and a target
	// it returns an error if the target cannot be made relative to the base
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
