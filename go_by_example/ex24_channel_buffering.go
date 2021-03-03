// by default channels are unbuffered and only accept sends with corresponding receives
// buffered channels accept a limited number of values without specifying a corresponding reciever

package main

import "fmt"

func main() {
	// here we make a channel of strings limited to 2 values
	messages := make(chan string, 2)

	// we can send values into unbuffered channels without corresponding recieves: val := <- channel
	messages <- "buffered"
	messages <- "channel"
	// because the limit is 2, we have to recieve a value before we can add one
	fmt.Println(<-messages)
	messages <- "three"

	// later we recieve those as usual
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
