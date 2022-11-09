// using Unix signals in go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// go signals work by sending os.Signal values on a channel
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// signal.Notify registers the channel to recieve specified signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// execute a blocking recieve for signals
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	// program waits untill it gets the expected signal, and then exits
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
