package main

import (
	"errors"
	"fmt"
)

// conventionally errors are the last return value, 
// and are type error, a built-in interface
func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New constructs an error value with the given message 
		return -1, errors.New("can't work with 42")
	}

	// nil in the error position if there is no error
	return arg + 3, nil
}

// by implementing the Error() method on a type
// you can use it as an error
type argError struct {
	arg int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		 
		// using & to build a new struct, and supply values for the two fields
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	// use of inline error check on the if line is
	// common idiom in Golang
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// to programmatically use the data in a custom error
	// you need to get the error as an instance of the error type
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

	ae2 := argError{2, "test2"}
	ae3 := &ae2
	ae3.arg = 3
	fmt.Println(*argError.Error(1, "test"))
	fmt.Println(ae2)
	fmt.Println(*ae3)
}

