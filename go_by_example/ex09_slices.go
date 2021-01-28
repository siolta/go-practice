package main

import "fmt"

func main() {

	// slices are type by the elements they contain, not the number
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	// slices support append
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// slices can also be copied
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// as expected, slices can be sliced
	l := s[2:5]
	fmt.Println("sl1:", l)

	// up to but excluding s[5]
	l = s[:5]
	fmt.Println("sl2:", l)

	// up from and including s[2]
	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// slices can be initialized shorthand:
	slice_nums := []int{1, 2, 3, 4, 5}
	fmt.Println("shorthand_nums:", slice_nums)

	// single and double quotes can't be mixed in a slice
	slice_strs := []string{"a", "b"}
	fmt.Println("slice_strs:", slice_strs)
}
