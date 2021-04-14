// examples of go's built-in regex support

package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	// test whether a pattern matches a string
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// for other tasks, use Compile to build a Regexp struct
	r, _ := regexp.Compile("p([a-z]+)ch")

	// this finds the match for the defined regexp
	fmt.Println(r.FindString("peach punch"))

	// this also finds the first match but returns the start and end indexes instead of the text
	fmt.Println(r.FindStringIndex("peach punch"))

	// submatch includes information about both the whole-pattern match
	// and the submatches, eg: p([a-z]+)ch and ([a-z]+)
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// likewise this returns the indexes of all matches and submatches
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// all finds all matches in the input, not just the first
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// All variants are available for the above functions
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	// providing a non-negative int limits the number of matches
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// can also use []byte args and drop String from the func name
	fmt.Println(r.Match([]byte("peach")))

	// MustCompile is safer to use for globals as it panics instead of error-ing
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	// regex can also replace subsets of strings with other values
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// the Func variant allows for transforming matched text with a given function
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
