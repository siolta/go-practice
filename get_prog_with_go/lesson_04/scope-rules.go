package main

import (
	f "fmt"
	r "math/rand"
)

var era = "AD" // era scoped to packaged

func main () {
	year := 2018 // era & year in scope for main func

	switch month := r.Intn(12) + 1; month { // era, year and month in scope for switch statement
	case 2:
		day := r.Intn(28) + 1 // era, year, month and day in scope for this case
		f.Println(era, year, month, day)
	case 4, 6, 9, 11:
		day := r.Intn(30) + 1 // it's a new day
		f.Println(era, year, month, day)
	default:
		day := r.Intn(31) + 1 // also a new day
		f.Println(era, year, month, day)
	} // month and day no longer in scope
} // year no longer in scope
