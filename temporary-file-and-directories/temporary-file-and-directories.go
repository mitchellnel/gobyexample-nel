package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Throughout program execution, we often want to create data that isn't needed after the program
// exits

// Temporary files and directories are useful for this purpose since they don't pollute the file
// system over time

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// the easiest way to create a temporary file is by calling os.CreateTemp
	// it creates a file and opens it for reading and writing

	// we will provide "" as the first argument, so os.CreateTemp will create the file in the
	// default location for our OS (likely /tmp for Unix)
	f, err := os.CreateTemp("", "sample")
	check(err)

	// display the name of the temporary file
	fmt.Println("Temp filename:", f.Name())
	// the filename starts with the prefix given as the second argument to os.CreateTemp, and the
	// rest is chosen automatically to ensure that concurrent calls will always create different
	// filenames

	// clean up the file after we're done
	defer os.Remove(f.Name())
	// the OS is likely to clean up temporary files by itself after some time, but it's good
	// practice to do this explicitly

	// we can write some data to the file
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	// if we intend to write many temporary files, we may prefer to create a temporary directory
	// os.MkdirTemp's arguments are the same as CreateTemp's, but it returns a directory name rather
	// than an open file
	dname, err := os.MkdirTemp("", "sampledir")
	check(err)
	fmt.Println("Temp directory name:", dname)

	defer os.RemoveAll(dname)

	// now we can synthesise temporary file names by prefixing them with our temporary directory
	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}
