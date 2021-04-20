// line filters like sed and grep can be implemented in go
// test data to filter: echo 'hello' > /tmp/lines && echo 'filter' >> /tmp/lines
// run with: cat /tmp/lines | go run ex59_line_filters.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// wrapping unbuffered os.Stdin gives a convenient Scan method
	scanner := bufio.NewScanner(os.Stdin)

	// Text returns the current token, which is the next line here
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())

		// write out uppercased line
		fmt.Println(ucl)
	}

	// check for errors during scan
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
