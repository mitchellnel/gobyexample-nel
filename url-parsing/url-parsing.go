package main

import (
	"fmt"
	"net"
	"net/url"
)

// URLs provide a uniform way to locate resources

// Here's how to parse URLs in Go

func main() {
	// we'll parse this example URL, which includes a scheme, authentication info, host, port, path,
	// query parameters, and query fragment
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// parse the URL and ensure that there are no errors
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// accessing the scheme is straightforward
	fmt.Println(u.Scheme)

	// User contains all authentication information
	// call Username and Password on this for individual values
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// the Host contains both the hostname and the port, if present
	// use SplitHostPort to extract them
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// here, we extract the path and the fragment after the #
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// to get query parameters in a string of k=v format, use RawQuery
	fmt.Println(u.RawQuery)

	// we can also parse query parameters into a map
	// the parsed query parameter maps are from strings to slices of strings, so index into [0] if
	// you only want the first value
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])

	// Running the above code will show all the different pieces of the URL that we extracted
}
