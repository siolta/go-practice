// implement worker pools by combining goroutines & channels

package main

import (
	"fmt"
	"time"
)

// worker implementation.  each worker recieves work on the jobs channel
// and sends the results to the results channel.  
// sleeping the second job simulates a long running task
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	// send work to the workers and collect results w/ 2 channels
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// start 3 workers, blocked initially as there are no jobs yet
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// send 5 jobs then close the channel to indicate that's all the work
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// collect all the results, also ensure all goroutines are done
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
