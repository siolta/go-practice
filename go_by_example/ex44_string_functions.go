// go's standard lib strings package examples

package main

import (
	"fmt"
	s "strings"
)

// alias this so it's easier to repeat below
var p = fmt.Println

// sampling of the functions in strings
// since these functions in strings, must pass the string as the first arg
func main() {
	p("Contains:  ", s.Contains("test", "es"))
    p("Count:     ", s.Count("test", "t"))
    p("HasPrefix: ", s.HasPrefix("test", "te"))
    p("HasSuffix: ", s.HasSuffix("test", "st"))
    p("Index:     ", s.Index("test", "e"))
    p("Join:      ", s.Join([]string{"a", "b"}, "-"))
    p("Repeat:    ", s.Repeat("a", 5))
    p("Replace:   ", s.Replace("foo", "o", "0", -1))
    p("Replace:   ", s.Replace("foo", "o", "0", 1))
    p("Split:     ", s.Split("a-b-c-d-e", "-"))
    p("ToLower:   ", s.ToLower("TEST"))
    p("ToUpper:   ", s.ToUpper("test"))
    p()

	// ways to get the length of a string in bytes, or getting a byte by index
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
}
