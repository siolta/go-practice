// Channels are pipes that connect goroutines
// You send values into channels from one goroutine
// and recieve them into another goroutine

package main

import "fmt"

func main() {
	// create a channel via make(chan val-type)
	// channels are typed by the values they convey
	messages := make(chan string, 2)

	// send a value to a channel via channel <- value syntax
	go func() { messages <- "ping"}()

	// <-channel syntax receives a value from the channel
	msg := <-messages
	fmt.Println(msg)
}

// By default sends and receives block until both the sender and receiver are ready.
// This property allowed us to wait at the end of our program for the "ping" message without having to use any other synchronization.
