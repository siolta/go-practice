// using flags instead of args in go
// best if built, then run as a binary
package main

// go provides a flag package for basic cli flag parsing
import (
	"flag"
	"fmt"
)

func main() {
	// flags are available for string, int, and bool options
	wordPtr := flag.String("word", "foo", "a string")

	// declare numb and fork flags
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	// declare an option that uses an existing var, need to pass in pointer
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// once all flags are declared, call flag.Parse()
	flag.Parse()

	// dump out parsed options and trailing positional args
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
