// examples of custom sorts in go

package main

import (
	"fmt"
	"sort"
)

// in order to custom sort, you need a corresponding type
// create a byLength type that is an ilias for the builtin []string type
type byLength []string

// implement Len, Less, and Swap on the byLength type
func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less implements the custom sorting logic
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
