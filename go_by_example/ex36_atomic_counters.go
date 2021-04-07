// channels are the primary method of managing state in go
// atomic counters are another possible method

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// use an unsigned int to represent the counter
	var ops uint64

	// use waitgroup so all goroutines finish
	var wg sync.WaitGroup

	// start 50 goroutines that increment the counter 1000 times
	for i := 0; i < 50; i++ {
		wg.Add(1)

		// increment the counter using AddUint64
		go func() {
			for c := 0; c < 1000; c ++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	// wait for all goroutines to finish
	wg.Wait()

	// safe to access ops now since we know no other routine is writing to it.
	// atomics can be safely read as they are updated using funcs like atomic.LoadUint64
	fmt.Println("ops:", ops)
}
