// if you want to wait for multiple goroutines to finish
// use a waitgroup

package main

import (
	"fmt"
	"sync"
	"time"
)

// this function will run in every goroutine
// note WaitGroup must be passed by pointer
func worker(id int, wg *sync.WaitGroup) {
	// on return notify the WaitGroup we're done
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)
	// sleep to simulate expensive task
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// this waitgroup is used to wait for all goroutines to finish
	var wg sync.WaitGroup

	// launch several goroutines and increment wg counter for each
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// block until the wg counter goes to 0
	wg.Wait()
}
