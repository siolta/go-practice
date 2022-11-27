package main

import (
	f "fmt"
	t "time"
)

func main() {
	f.Println("7.1:")
	inspect()
	f.Println("7.2:")
	wrap()
	f.Println("7.3:")
	bits()
	f.Println("7.5:")
	time()
}

func inspect() {
	year := 2018
	f.Printf("Type %T for %v\n", year, year)

	// can also uso [1] syntax to avoid repeating the var twice
	days := 365.2425
	f.Printf("Type %T for %[1]v\n", days)

	//Qc 7.3
	f.Printf("Type %T for %[1]v\n", "text")
	f.Printf("Type %T for %[1]v\n", 42)
	f.Printf("Type %T for %[1]v\n", 3.14)
	f.Printf("Type %T for %[1]v\n", true)
}

func wrap() {
	// ints wrap around when they increment beyond allowed range
	var red uint8 = 255
	red++
	f.Println(red)

	var number int8 = 127
	number++
	f.Println(number)
}

func bits() {
	// us %b format verb to show bits for int values
	var green uint8 = 3
	f.Printf("%08b\n", green)
	green++
	f.Printf("%08b\n", green)
	// 7.4
	var blue uint8 = 255
	f.Printf("%08b\n", blue)
	blue++
	f.Printf("%08b\n", blue)
}

func time() {
	future := t.Unix(12622780800, 0)
	f.Println(future)
}
