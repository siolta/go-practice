package main

import "fmt"
// Variadic functions can be called with any number of arguments
// fmt.Println is a common variadic function

// ... makes a func take an arbitrary number of args
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	// variadic funcs can be called as usual
	sum(1, 2)
	sum(1, 2, 3)

	// or if you have a slice, can be called as so
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
