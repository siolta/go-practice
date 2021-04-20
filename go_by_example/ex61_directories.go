// functions for working with directiories in go
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"pathe/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// create new subdirectory in current working dir
	err := os.Mkdir("subdir", 0755)
	check(err)

	// when making temp dirs, defer their removal
	defer os.RemoveAll("subdir")

	// helper func to create empty file
	createEmptyFile := func(name string) {
		d := []byte("")
		check(ioutil.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	// create heirarchies like with mkdir -p
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// ReadDir lists directory contents, returs slice of os.FileInfo objects
	c, err := ioutil.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// Chdir changes directory like cd
	err = os.Chdir("subdir/parent/child")
	check(err)

	// see contents in current dir
	c, err = ioutil.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// cd back to start
	err = os.Chdir("../../..")
	check(err)

	// Walk takes a callback to recursively handle each dir listed
	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit)
}

// visit gets called for ever dir found by filepath.Walk
func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}
