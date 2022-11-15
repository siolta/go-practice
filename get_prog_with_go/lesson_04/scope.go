package main

import (
	f "fmt"
	r "math/rand"
)

func main() { // begins the first scope
	count := 0

	for count < 10 { // begins a second scope
		num := r.Intn(10) + 1
		f.Println(num)

		count++
	} // ends the second scope

	// 4.2 - Condensed countdown
	count = 0
	for count = 10; count > 0; count-- {
		f.Println(count)
	}
	f.Println(count) // count is still in scope since it was declared outside the loop

	// 4.3 - short var declaration in for loop
	for f_count := 10; f_count > 0; f_count-- {
		f.Println(f_count)
	} // for loop 'count' is no longer in scope

	// 4.4 - short declaration in if statement
	if num := r.Intn(3); num == 0 {
		f.Println("Space Adventures")
	} else if num == 1 {
		f.Println("SpaceX")
	} else {
		f.Println("Virgin Galactic")
	}

	// 4.5 - Short declaration in switch statement
	switch num := r.Intn(10); num {
	case 0:
		f.Println("Space Adventures")
	case 1:
		f.Println("SpaceX")
	case 2:
		f.Println("Virgin Galactic")
	default:
		f.Println("Random spaceline #", num)
	}
}
