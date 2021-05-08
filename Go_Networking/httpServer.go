package main

import (
	"fmt"
	"log"
	"net/http"
)

// helloHandler returns greeting

// A handler function is a function that gets an http.ResponseWriter where you write the response to
//and a pointer to an http.Request, which represents the current request.
//Since the http.ResponseWriter implements io.writer, we can use the Fprintf to print the output to it.
// We need to mount the endpoint to the Handler. This means that every time someone accesses /hello on our web server,
//helloHandler is going to get called. And now we tell the HTTP server to listen and serve on port 8080.
// This is potentially an endless loop. But if there is an error, it will exit with an error value.
//And in our case, we're going to print out the error and exit the program.

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello there go aspirants !!")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
