// basic HTTP server using net/http package
package main

import (
	"fmt"
	"net/http"
)

// http handler function
func hello(w http.ResponseWriter, req *http.Request) {
	// handler functions take a http.ResponseWriter, and a http.Request as args
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	// this handler returns all the HTTP request headers and echoing them into the response body
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	// register handlers on server routes using http.HandleFunc
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// call the ListenAndServe function, with port and handler, nil uses the default router
	http.ListenAndServe(":8090", nil)
}
