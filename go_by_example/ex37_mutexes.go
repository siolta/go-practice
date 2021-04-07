// mutexes can be used to manage more complex state across multiple goroutines

 package main
 import (
	 "fmt"
	 "math/rand"
	 "sync"
	 "sync/atomic"
	 "time"
 )

 func main() {
	 var state = make(map[int]int)

	 // this mutex will synchronize access to 'state'
	 var mutex = &sync.Mutex{}

	 var readOps uint64
	 var writeOps uint64

	 // start 100 goroutines to simulate reads
	 for r := 0; r < 100; r++ {
		 go func() {
			 total := 0
			 for {
				 // for each read, pick a key to access
				 key := rand.Intn(5)
				 // Lock() the mutex to ensure exclusive access
				 mutex.Lock()
				 // read the value
				 total += state[key]
				 // Unlock() the mutex
				 mutex.Unlock()
				 // increment the readOps count
				 atomic.AddUint64(&readOps, 1)

				 time.Sleep(time.Millisecond)
			 }
		 }()
	 }

	 // start 10 goroutines to simulate writes
	 for w := 0; w < 10; w++ {
		 go func() {
			 for {
				 key := rand.Intn(5)
				 val := rand.Intn(100)
				 mutex.Lock()
				 state[key] = val
				 mutex.Unlock()
				 atomic.AddUint64(&writeOps, 1)
				 time.Sleep(time.Millisecond)
			 }
		 }()
	 }

	 // let the goroutines work for a second
	 time.Sleep(time.Second)

	 // total final operation counts
	 readOpsFinal := atomic.LoadUint64(&readOps)
	 fmt.Println("readOps:", readOpsFinal)
	 writeOpsFinal := atomic.LoadUint64(&writeOps)
	 fmt.Println("writeOps:", writeOpsFinal)

	 // final lock of the state
	 mutex.Lock()
	 fmt.Println("state", state)
	 mutex.Unlock()
 }
