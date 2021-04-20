// reading and writing files in go
// put some data in file first with: echo "hello" > /tmp/dat && echo "go" >> /tmp/dat
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// reading files requires checking for errors, this is a helper func to do so
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// slurp files contents into memory
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	// often to control what parts of a file are read you need to use Open to get an os.File value
	f, err := os.Open("/tmp/dat")
	check(err)

	// read up to 5 bytes from beginning of file, and note how many are read
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// Seek to known location and Read from there
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// io provides functions that are helpful for reading files
	// more robust version of above read
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// no built-in rewind, but Seek(0, 0) accomplishes this
	_, err = f.Seek(0, 0)
	check(err)

	// bufio implements a buffered reader that is useful for efficiency and additional read methods
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// close file when done, usually scheduled right after Open with defer
	f.Close()
}
