package main

import (
	f "fmt"
	r "math/rand"
)

const (
	distanceToMars = 62100000
	secondsPerDay = 86400
)

func main() {
	f.Printf("%-15v %4v %-10v %4v\n", "Spaceline", "Days", "Trip type", "Price")
	f.Println("=====================================")
	for c := 10; c > 0; c-- {
		f.Printf("%-15v %4v %-10v $%4v\n", company(), days(), tripType(), cost())
	}
}

func days() int {
	speed := r.Intn(15) + 16 // 16-30 km/s
	return distanceToMars / speed / secondsPerDay 
}

func company() string {
	if c := r.Intn(2); c == 0 {
		return "SpaceX"
	} else if c == 1 {
		return "Virgin Galactic"
	} else {
		return "Space Adventures"
	}
}

func tripType() string {
	if t := r.Intn(2); t == 1 {
		return "Round-trip"
	} else {
		return "One-way"
	}
}

func cost() int {
	return r.Intn(100)
}
