// go supports rate limiting with goroutines, channels, and tickers

package main

import (
	"fmt"
	"time"
)

func main() {
	// limit handling of incoming requests
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// limiter recieves a value every 200ms
	limiter := time.Tick(200 * time.Millisecond)

	// by blocking on the limiter channel we limit to 1 req per 200ms
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// this limiter allows bursts of up to 3 requests
	burstyLimiter := make(chan time.Time, 3)

	// fill the bursty channel to represent a burst
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// add a value to bursty every 200ms
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// simulate 5 more incoming reqs, first 3 bursty
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
