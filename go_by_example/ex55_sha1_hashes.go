// compute sha1 hashes in go
// go has several hash functions in crypto/*
package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "sha1 this string"

	// to generate a hash, sha1.New, sha1.Write(bytes), then sha1.Sum([]byte{})
	h := sha1.New()

	// write expects bytes, use []byte(s) to coerce it to byte
	h.Write([]byte(s))

	// get the final result as a byte slice, arg can be used to append, but usually isn't needed
	bs := h.Sum(nil)

	// use %x to convert result to hex string
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
