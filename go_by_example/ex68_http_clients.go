// http requests in go
package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	// issue HTTP GET to server, http.get is a shortcut to creating a client and calling its get method
	resp, err := http.Get("http://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// print HTTP response status
	fmt.Println("Response status:", resp.Status)

	// print first 5 lines of response
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
