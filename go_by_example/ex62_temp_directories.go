// creating and deleting temp dirs in go
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// ioutil.TempFile creates and opens file for rw, first arg is location
	f, err := ioutil.TempFile("", "sample")
	check(err)

	// display the name of the temp file
	fmt.Println("Temp file name:", f.Name())

	// clean up file after we're done
	defer os.Remove(f.Name())

	// write data to file
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	// if intending to write many temp files, use TempDir
	dname, err := ioutil.TempDir("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	// sythisize temp files by prefixing them with the temp dir
	fname := filepath.Join(dname, "file1")
	err = ioutil.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}
