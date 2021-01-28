package main

import "fmt"

func main() {
	// map takes the key-type & the value-type
	// make(map[key_type]val_type)
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	// getting a value has an optional second return value that indicates if the value exists
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// declare and initialize a new map
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
