package main

import (
	f "fmt"
)

func main() {
	message := "L fdph L vdz L frqtxhuhg."

	for i := 0; i < len(message); i++ {
		c := message[i]
		if c >= 'a' && c <= 'z' {
			c -= 3
			if c > 'z' {
				c = c - 26
			}
		} else if c >= 'A' && c <= 'Z' {
			c -= 3
			if c > 'Z' {
				c = c - 26
			}
		}
		f.Printf("%c", c)
	}
}
