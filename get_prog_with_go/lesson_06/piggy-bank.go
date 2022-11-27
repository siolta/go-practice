package main

import (
	f "fmt"
	r "math/rand"
)

func main () {
	for b := 0.0; b < 20.0; {
		switch r.Intn(3) {
		case 0:
			b += 0.05 
		case 1:
			b += 0.10
		case 2:
			b += 0.25
		}
		f.Printf("$%4.2f\n", b)
	}
}
