// go's filepath package has os agnostic filepath constructors
package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	// Join constructs paths in a portable way
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p: ", p)

	// prefer Join over concatenating as it normalizes paths
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// Dir and Base can be used to split a path
	fmt.Println("Dir(p): ", filepath.Dir(p))
	fmt.Println("Base(p): ", filepath.Base(p))

	// check if a path is absolute
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	// some file names have extensions, split using Ext
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// to get the name without the extension, use strings.TrimSuffix
	fmt.Println(strings.TrimSuffix(filename, ext))

	// Rel finds a relative path between a base and target, returns error if target cannot be made relative
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
