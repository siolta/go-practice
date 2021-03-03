package main

import (
	"fmt"
	"time"
)

// goroutines are lightweight execution threads

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// call a function syncronously
	f("direct")

	// invoke in a goroutine, will execute concurrently
	go f("goroutine")

	// can also start a goroutine for an anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
