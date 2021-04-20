// writing files follows a similar pattern to reading
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// how to dump a string or bytes into a file
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// for more granular writes, open file for writing
	f, err := os.Create("/tmp/dat2")
	check(err)

	// idiomatic to defer a Close immediately after opening a file
	defer f.Close()

	// can Write byte slices as expected
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// WriteString is also available
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// Sync to flush writes to stable storage
	f.Sync()

	// bufio also has buffered writes
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	// Flush to ensure all buffered operations have been applied
	w.Flush()
}
