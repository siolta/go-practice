package main

import (
	f "fmt"
	m "math"
)

func main () {
	third := 1.0 / 3

	// %f can be used with float* to specify number of digits as follows:
	// %{width}.{precision}f : width == min chars to display precision == digits after decimal
	f.Println(third) // print as many as possible
	f.Printf("%v\n", third) // print as many as possible
	f.Printf("%f\n", third) // print  0.333333
	f.Printf("%.3f\n", third) // print  0.333
	f.Printf("%4.2f\n", third) // print  0.33
	// zero padding
	f.Printf("%06.2f\n", third) // print  0.33

	// rounding errors:
	f.Println(third * 3)
	piggyBank := 0.1
	piggyBank += 0.2
	// can use precision of 2 digits to hide rounding errors
	f.Println(piggyBank)
	// rounding error
	celsius := 21.0
	f.Print((celsius/5.0*9.0)+32, "° F\n")
	// rounding errors can be multiplying first
	fahrenheit := (celsius * 9.0 / 5.0) + 32.0
	f.Print(fahrenheit, "° F\n")
	// floats when compared directly can be unintuitive due to rounding errors
	f.Println(piggyBank == 0.3)
	// instead compare the difference between them and ensure it isn't too big
	f.Println(m.Abs(piggyBank-0.3) < 0.0001)

	piggyBank = 0.0
	for i := 0; i < 11; i++ {
		piggyBank += 0.1
	}
	f.Println(piggyBank)
}
