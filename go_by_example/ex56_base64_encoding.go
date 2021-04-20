// go has built in base64 encoding/decoding support
package main

import (
	// this syntax imports the package with the preceding name
	b64 "encoding/base64"
	"fmt"
)

func main() {
	// string to encode / decode
	data := "abc123!?$*&()'-=@~"

	// go supports standard and URL-compatible base64, using the standard encoder
	// requires bytes, so we convert with []byte
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// decoding can return an error, so we check
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// encode/decode using URL compatible base64
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
