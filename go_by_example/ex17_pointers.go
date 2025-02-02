// pointers let you pass references to values

package main

import "fmt"

// zeroval has arguments passed by value
func zeroval(ival int) {
	ival = 0
}

// zeroptr uses pointer parameters via the * syntax
func zeroptr(iptr *int) {
	// *iptr derefences the pointer from its memory address
	// to the current value at that address
	// assigning a value to a d'ref'd pointer changes the value
	// at the referenced address
	// aka '*' means 'value at' the referenced address
	*iptr = 0
}

func main() {
	i := 1
	b := 3
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// the & operator gives the memory 'address of' i
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// pointers can also be printed
	fmt.Println("pointer:", &i)

	fmt.Println("pointer2:", &b)
}
