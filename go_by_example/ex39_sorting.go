// go's sort package handles sorting for builtins and user defined types
package main

import (
	"fmt"
	"sort"
)

func main() {
	// sort methods are specific to the builtin type, this is for strings
	// note it sorts in-place
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// sorting ints
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:  ", ints)

	// can use sort to check if a slice is already sorted
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
