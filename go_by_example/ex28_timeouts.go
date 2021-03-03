// timeouts can be constructed using channels and select
package main

import (
	"fmt"
	"time"
)

func main() {
	// simulate external call that returns result after 2s
	// channel is buffered so the send isn't blocking
	// commonly used to prevent leaks in case the channel is never read
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// select as a timeout
	select {
	// wait for the result
	case res := <-c1:
		fmt.Println(res)
	// sends a value after one second(the timeout)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// set the timeout to 3s so the reciever finishes
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
