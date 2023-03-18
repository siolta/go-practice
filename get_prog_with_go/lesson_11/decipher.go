package main

import (
	f "fmt"
)

func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"

	for i := 0; i < len(cipherText); i++ {
		if len(keyword) < len(cipherText) {
			keyword += keyword
		}
		c := cipherText[i]
		k := keyword[i]
		f.Printf("%c", c)
	}
}
