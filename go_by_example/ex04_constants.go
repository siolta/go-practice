package main

import (
	"fmt"
	"math"
)

// const declares constant values
const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	// numeric constants have no type until given one
	fmt.Println(int64(d))

	// unless given a type by its context
	fmt.Println(math.Sin(n))
}
