package main

import (
	"fmt"
	"net/http"
	"time"
)

// HTTP servers are useful for demonstrating the usage of context.Context for controlling
// cancellation

// A Context carries deadlines, cancellation signals, and other request-scoped values across API
// boundaries and goroutines

func hello(w http.ResponseWriter, req *http.Request) {
	// a context.Context is created for each request by the net/http machinery, and is available
	// with the Context() method
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler finished")

	// wait a few seconds before sending a reply to the client
	// this could simualte some work that the serer is doing
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		// while working, keep an eye on the context's Done() channel for a signal that we should
		// cancel the work and return as soon as possible

		// the context's Err() methods returns an error tha explains why the Done() channel was
		// closed
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {
	// register our handler and start serving
	http.HandleFunc("/hello", hello)

	err := http.ListenAndServe(":8091", nil)
	if err != nil {
		panic(err)
	}

	// Run the server
	// >> go run . &

	// Simulate a client request to /hello, hitting Ctrl+C shortly after starting to send a SIGINT
	// >> curl localhost:8091/hello
	// >> ^C
}
