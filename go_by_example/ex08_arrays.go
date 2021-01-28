package main

import "fmt"

func main() {
	// arrays in Go are numbered sequences of elements of a given length
	// the type of elements, and length are both part of the array's type
	// zero-valued by default
	var a [5]int
	fmt.Println("emp:", a)


	// set values like usual
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	// declare and initialize in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
