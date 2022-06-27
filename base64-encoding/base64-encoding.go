package main

// Go provides built-in support for base64 encoding/decoding

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	// here's the string that we'll encode/decode
	data := "abc 123!?$*&()'-=@~"

	// Go supports both standard and URL-compatible base64

	// here's how to encode using the standard encoder
	// the encode requires a []byte, so we convert our string to that type
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// decoding may return an error, which you can check if you don't already know the input to be
	// well-formed
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// this encodes/decodes using a URL-compatible base 64 format
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
