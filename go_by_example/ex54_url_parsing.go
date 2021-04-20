// url parsing in go
package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
		// example URL to parse
		s := "postrgres://user:pass@host.com:5432/path?k=v#f"

		// parse url and check for errors
		u, err := url.Parse(s)
		if err != nil {
			panic(err)
		}

		// get the scheme
		fmt.Println(u.Scheme)

		// user and password
		fmt.Println(u.User)
		fmt.Println(u.User.Username())
		p, _ := u.User.Password()
		fmt.Println(p)

		// Host contains both the hostname and port, SplitHostPort extracts them
		fmt.Println(u.Host)
		host, port, _ := net.SplitHostPort(u.Host)
		fmt.Println(host)
		fmt.Println(port)

		// get path and fragments after the #
		fmt.Println(u.Path)
		fmt.Println(u.Fragment)

		// use RawQuery to get params of k=v format
		fmt.Println(u.RawQuery)
		m, _ := url.ParseQuery(u.RawQuery)
		fmt.Println(m)
		fmt.Println(m["k"][0])
}
