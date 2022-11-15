package main

import (
	f "fmt"
	r "math/rand"
)

func main() {
	r.Seed(50)
	number := r.Intn(99) + 1

	f.Printf("Picked number is %v\n", number)

	for {
		guess := r.Intn(99) + 1

		if guess == number {
			f.Printf("Success! The number you chose is: %v\n", guess)
			break
		} else if guess > number {
			f.Printf("%v was too high.\n", guess)
		} else if guess < number {
			f.Printf("%v was too low.\n", guess)
		}
	}
}
