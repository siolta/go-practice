// set, get and list env vars in go
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// to set k/v pair use os.Setenv, retrieve w/ os.Getenv
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	// use os.Environ to list all k/v pairs
	// returns a slice of strings in the form KEY=value
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
