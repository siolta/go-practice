package main

import "fmt"

func main() {

	// var declares 1 or more variables
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)
	
	// go infers type of initialized variables
	var d = true
	fmt.Println(d)

	// variables without value, are zero-valued
	var e int
	fmt.Println(e)

	// := shorthand for declare, and initialize
	f := "apple"
	fmt.Println(f)
}
