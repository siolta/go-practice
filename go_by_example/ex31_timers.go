// timers can be used to execute code repeatedly, or in the future
package main

import (
	"fmt"
	"time"
)

func main() {
	// you specify how long to wait, and it sends it's channel at that time
	timer1 := time.NewTimer(2 * time.Second)

	// block on the timers C channel until the value is sent
	<-timer1.C
	fmt.Println("timer 1 fired")

	// you can cancel timers before they fire
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}
