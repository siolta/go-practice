// go uses 'contexts' for deadlines, cancellation signals, and other request-scoped values
// simulate request by curling localhost, then cancelling with ctl-c
package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	// a context.Context is created for each request in net/http
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	// wait a few seconds before sending reply and watch the ctx.Done() channel for done signal
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():

		// contexts Err() method returns an error that explains why Done() was closed
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {
	// register handler on /hello route
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
