package main

import "fmt"

// int, int declares the return type of the function
func vals() (int, int) {
	return 3, 7
}

func main() {
	// go multiple return value assignment
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// _ is commonly used to signify a value isn't needed
	_, c := vals()
	fmt.Println(c)
}
