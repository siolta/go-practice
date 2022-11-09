// os.Exit() imediately exits with a given status
package main

import (
	"fmt"
	"os"
)

func main() {
	// defers will _not_ run when using os.Exit()
	defer fmt.Println("!")
	
	// exit with status 3
	os.Exit(3)
}
