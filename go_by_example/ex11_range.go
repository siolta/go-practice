package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	// range returns the value, and the index
	// _ is used to discard the index
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// range on map iterates over k/v pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s/n", k, v)
	}

	// or it can iterate over just the keys
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range on strings iterates over unicode code points
	// i is the byte index, and c is the rune
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
