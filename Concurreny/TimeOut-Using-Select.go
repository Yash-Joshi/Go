package main

import (
	"fmt"
	"time"
)

func main() {

	out := make(chan float64)

	go func() {
		time.Sleep(100 * time.Millisecond)
		out <- 3.14
	}()

	select {
	case val := <-out:
		fmt.Printf("got %f\n", val)

	case <-time.After(20 * time.Millisecond):
		fmt.Printf("timeout")
	}
}

// Summary
// using case <-time.After(20 * time.Millisecond):
// will output timeout
// but this does not kill go routine in behind.

// using case <-time.After(20 * time.Millisecond):
// will print the value receieved on the channel

// note: that you can use default channel to in this
