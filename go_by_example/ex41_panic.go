// panics are used when something goes unexpectedly wrong
// mostly used to fail fast on errors that shouldn't occur during normal operation

package main

import "os"

func main() {
	panic("a problem")

	// common use of panic to abort if a func returns an error we don't know how to handle
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
