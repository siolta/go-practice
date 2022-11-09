// best if built, then run as a binary
package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args gives access to raw cli args, first val is the path to the program
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// get individual args by indexing normally
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
