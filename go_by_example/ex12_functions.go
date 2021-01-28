package main

import "fmt"

//func_name(param_a a_type, param_b type_b) return type 
func plus(a int, b int) int {
	// go requires explicit returns
	return a + b
}

// you can group together multiple parameters of the same type
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	// call functions as usual, with name(args)
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	// you can skip assigning the result
	fmt.Println("2+3 =", plus(2, 3))
}
