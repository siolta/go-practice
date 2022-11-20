package main

import (
	f "fmt"
	// r "math/rand"
)

func main () {
	third := 1.0 / 3

	// %f can be used with float* to specify number of digits as follows:
	// %{width}.{precision}f : width == min chars to display precision == digits after decimal
	f.Println(third) // print as many as possible
	f.Printf("%v\n", third) // print as many as possible
	f.Printf("%f\n", third) // print  0.333333
	f.Printf("%.3f\n", third) // print  0.333
	f.Printf("%7.2f\n", third) // print  0.33
}
