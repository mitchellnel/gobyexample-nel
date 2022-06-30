package main

import (
	"fmt"
	"net/http"
)

// Writing a basic HTTP server is easy using the net/http package

// A fundamental concept in net/http servers are **handlers**
// A handler is an object implemeneting the http.Handler interface
// A common way to write a handler is using the http.HandlerFunc adapter on functions with the
// appropriate signature

func hello(w http.ResponseWriter, req *http.Request) {
	// function serving as handlers take a http.ResponseWrite and a http.Request as arguments
	// the respone writer is used to fill in the HTTP response

	// here, our simple response is just "hello\n"
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	// this handler does something a little more sophisticaed by reading all the HTTP request
	// headers and echoing them into the response body

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	// we register our handlers on server routes using the http.HandleFunc convenience functions
	// it sets up the default router in the net/http package and takes a function as an argument
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// finally we call the ListenAndServe with the port and a handler
	// nil tells it to use the default router we've just set up
	http.ListenAndServe(":8090", nil)

	// Then, to run the server
	// >> go run . &

	// Then access the /hello route
	// >> curl localhost:8090/hello
}
