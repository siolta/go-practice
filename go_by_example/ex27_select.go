// select lets you wait on multiple channels

package main

import (
	"fmt"
	"time"
)

func main() {
	// select across two channels
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan int)

	// each channel gets a value after som time passes
	// this simulates blocking RPC operations
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()
	go func() {
		time.Sleep(3 * time.Second)
		c3 <- 3
	}()

	// loop over select to await all values simultaneously-ish
	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("recieved_1", msg1)
		case msg2 := <-c2:
			fmt.Println("recieved_2", msg2)
		case msg3 := <-c3:
			fmt.Println("recieved_3", msg3)
		}
	}
}
