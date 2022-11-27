package main
// piggy with ints instead of floats

import (
	f "fmt"
	r "math/rand"
)

func main () {
	for b := 0; b < 2000; {
		switch r.Intn(3) {
		case 0:
			b += 5 
		case 1:
			b += 10
		case 2:
			b += 25
		}
		dollars := b / 100
		cents := b % 100
		f.Printf("$%d.%02d\n", dollars, cents)
	}
}
