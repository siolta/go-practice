package main

import "fmt"

func intSeq() func() int {
	i := 0
	// this func is defined anonymously and closes over the var i
	return func() int {
		i++
		return i
	}
}

func main() {

	// calling intSeq captures and returns an updated i on each call
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// i's state is unique to each instance
	newInts := intSeq()
	fmt.Println(newInts())
}
