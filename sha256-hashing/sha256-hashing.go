package main

import (
	"crypto/sha256"
	"fmt"
)

// SHA256 hashes are frequently used to compute short identities for binary or text blobs

// For example, TLS/SSL certificates use SHA256 to compute a certificate's signature

// Here's how we compute SHA256 hashes in Go

// Go implements several hash functions in various crypto/* packages

func main() {
	s := "sha256 this string"

	// here we start with a new hash
	h := sha256.New()

	// Write expects bytes
	// so if we have string s, use []byte(s) to coerce it to a slice of bytes
	h.Write([]byte(s))

	// this gets the finalised hash result as a byte slice
	// the argument to Sum can be used to append to an existing byte slice -- it usually isn't
	// needed
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)

	// Running the program computes the hash and prints it in a human-readable hex format

	// You can compute other hashes using a similar pattern to the one shown above
	// For example, to compute SHA512 hashes, we import crypto/sha512 and use sha512.New()
}
