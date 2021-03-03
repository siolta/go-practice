// tickers are used when you want to repeat something at regular intervals

package main

import (
	"fmt"
	"time"
)

func main() {
	// tickers use similar mechinisms to timers
	// a channel is created that values are sent to
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// tickers can be stopped just like timers
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
