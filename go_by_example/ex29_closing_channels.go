// closing a channel indicates that no more values will be send on it
package main

import "fmt"

func main() {
	// jobs tracks work to be done
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			// more will == false if jobs has been closed and all values received
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// send 3 jobs to the 'worker', then close the channel
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// wait for the worker using sychronization with a blocking recieve
	<-done
}
