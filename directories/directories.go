package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Go has several useful functions for working with directories in the file system

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// create a new sub-directory in the current working directory
	err := os.Mkdir("subdir", 0755)
	check(err)

	// when creating temporary directories, it's good practice to defer their removal
	// os.RemoveAll will delete a whole directoy tree (similarly to rm -rf)
	defer os.RemoveAll("subdir")

	// helper function to create a new empty file
	createEmptyFile := func(name string) {
		d := []byte("")

		err := os.WriteFile(name, d, 0644)
		check(err)
	}

	createEmptyFile("subdir/file1")

	// we can createa hierarchy of directories including parents with MkdirAll
	// this is similar to the command line mkdir -p
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// ReadDir lists directory contents, returning a slice of os.DirEntry objects
	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// Chdir lets us change the current working directory, similary to cd
	err = os.Chdir("subdir/parent/child")
	check(err)

	// now we'll see the contents of subdir/parent/child when listing the current directory
	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// cd back to where we started
	err = os.Chdir("../../..")
	check(err)

	// we can also visit a directory recursively, including all its sub-directories
	// Walk accepts a callback function to handle every file or directory visited
	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit)
	check(err)
}

// visit is called for every file or directory found recursively by filepath.Walk
func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}
