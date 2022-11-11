package main

import (
	f "fmt"
	r "math/rand"
)

func main() {
	r.Seed(57) // initializes default 'source', if not called generator behaves as if 'Seed(1)'
	var num = r.Intn(10) + 1 // := is shorthand for variable creation and assignment
	f.Println(num)

	num = r.Intn(10) + 1
	f.Println(num)

	distance := r.Intn(401000001) - 56000000
	f.Println(distance)
}
