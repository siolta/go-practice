// You can also use the built in synchronization features of goroutines
// to manage state across multiple goroutines

package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// these structs encapsulate requests and provide a way for the goroutine to respond
type readOp struct {
	key int
	resp chan int
}
type writeOp struct {
	key int
	val int
	resp chan bool
}

func main() {
	// count the ops
	var readOps uint64
	var writeOps uint64

	// channels to be used by other goroutines to read and write requests
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// this goroutine selects on the reads and writes channels
	// responding to requests as they arrive
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// start goroutines to issue reads to the above goroutine via the reads channel
	for r := 0; r < 100; r++ {
		go func() {
			for {
				// construct readop
				read := readOp{
					key: rand.Intn(5),
					resp: make(chan int)}
				// send it over the reads channel
				reads <- read
				// recieve result via resp channel
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// start 10 writes using same pattern
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key: rand.Intn(5),
					val: rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// let the routines work for a second
	time.Sleep(time.Second)

	// collect and report results
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps", writeOpsFinal)
}
