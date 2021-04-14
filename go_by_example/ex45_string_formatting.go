// examples of common string formatting

package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	// print an instance of the point struct
	fmt.Printf("%v\n", p)

	// if value is a struct %+v includes the field names
	fmt.Printf("%+v\n", p)

	// print syntax representation of the value, IE the code that produced the value
	fmt.Printf("%#v\n", p)

	// print the type of a value
	fmt.Printf("%T\n", p)

	// format a boolean
	fmt.Printf("%t\n", p)
	
	// %d gives standard base10 int formatting
	fmt.Printf("%d\n", 123)

	// print a binary representation
	fmt.Printf("%b\n", 14)

	// print the character that corresponds to the given int
	fmt.Printf("%c\n", 33)

	// hex encoding
	fmt.Printf("%x\n", 456)

	// base decimal float formatting
	fmt.Printf("%f\n", 78.9)
	
	// two variants of scientific notation
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	// basic string formatting
	fmt.Printf("%s\n", "\"string\"")

	// double quote strings
	fmt.Printf("%q\n", "\"string\"")

	// render string in base-16
	fmt.Printf("%x\n", "hex this")

	// representation of a pointer
	fmt.Printf("%p\n", &p)

	// right justify and specify width/precision for ints
	fmt.Printf("|%6d|%6d|\n", 12, 345)

	// for floats
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	// left justify with '-' flag
	fmt.Printf("|%-6.2f|%-6.2f|\n", 12, 345)

	// table like output for strings
	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	// left justify same as with ints, '-' flag
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	// Sprintf formats and returns the string, without printing it
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	// format+print to io.Writers other than os.Stdout using Fprintf
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
