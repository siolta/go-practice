// Arrays

package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	a[0] = 10
	fmt.Println("set:", a)
	fmt.Println("get:", a[0])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl1: ", b)

	c := [2]string{"str1", "str2"}
	fmt.Println("dcl2: ", c)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	var threeD [3][3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				threeD[i][j][k] = i + j + k
			}
		}
	}
	fmt.Println("3d: ", threeD)
}
