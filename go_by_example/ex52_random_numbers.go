// random number generation in go with math/rand
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// rand.Intn returns a random int n 0 <= n < 100
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	// rand.Float64 returns a float64 f, 0.0 <= f < 1.0
	fmt.Println(rand.Float64())

	// can use this to generate random floats in other ranges
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// default generator is deterministic, to get varying sequences
	// give it a seed that changes
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// call like before
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	// using the same seed produces the same sequence of numbers
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	// s3 := rand.NewSource(42)
	r3 := rand.New(rand.NewSource(42))
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}
