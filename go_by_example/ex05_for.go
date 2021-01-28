package main

import "fmt"

func main() {

	// for is the only loop in go
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// initial/condition/after for loop
	for j := 7; j <=9; j++ {
		fmt.Println(j)
	}

	// while equivalent
	for {
		fmt.Println("loop foreeeverr")
		break
	}	

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
