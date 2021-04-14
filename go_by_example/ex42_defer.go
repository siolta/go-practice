// defer is used in go to ensure a function call is performed later
// in a programs execution, usually for cleanup
// often used where `ensure` or `finally` would be used

package main

import (
	"fmt"
	"os"
)

func main() {
	// create file, write, and then close when done using defer
	f := createFile("/tmp/defer.txt")
	// defer executes the declared function at the end of the enclosing function (main)
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	// important to check for errors when closing a file
	err := f.Close()

	if err != nil {
		fmt.Fprintln(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
