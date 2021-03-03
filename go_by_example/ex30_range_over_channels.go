// for and range can be used to interate over values received from a channel

package main
import "fmt"

func main() {
	// interate over two values in the queue channel
	queue := make(chan string, 3)
	queue <- "one"
	queue <- "two"
	// channels don't have to be empty when they are closed
	close(queue)

	// range iterates over elements in queue, 2 in this case since the channel is closed
	for elem := range queue {
		fmt.Println(elem)
	}
}
