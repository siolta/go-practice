package main

import (
	f "fmt"
	r "math/rand"
)

var era = "AD"

func main() {
	r.Seed(50)
	// ten random dates:
	for count := 10; count > 0; count-- {
		date()
	}
}

func date() {
	year := r.Intn(2022) + 1
	leap := year%400 == 0 || (year%4 == 0 && year%100 !=0)
	month := r.Intn(12) + 1
	daysInMonth := 31

	switch month {
	case 2: // case introduces a new scope even w/o braces
	daysInMonth = 28
		if leap {
			daysInMonth = 29
		}
	case 4, 6, 9, 11:
		daysInMonth = 30
	}

	day := r.Intn(daysInMonth) + 1
	f.Println(era, year, month, day)
}
