// channels can be used to synch execution across goroutines

package main

import (
	"fmt"
	"time"
)

// this will run in a goroutine
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// the done channel will be used to notify another
	// goroutine that this func is done
	done <- true
}

func main() {
	// start the worker go routine, supplying the channel to notify
	done := make(chan bool, 1)
	go worker(done)

	// block until the notification is received
	// if removed the program will exit before the worker starts
	<-done
}
