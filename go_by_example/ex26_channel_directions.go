// when used as params you can specify if a channel is send or recieve only
package main

import "fmt"

// ping only accepts a send only channel, will error at compile if you try to recieve
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pong accepts one channel for receive (pings) and another for sends (pongs)
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
