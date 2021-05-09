package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		// Pushing number to channel
		ch <- 888

	}()

	// here we are receiving the value from the channel
	val := <-ch

	start := time.Now()
	fmt.Printf("got %d\n", val)

	fmt.Println("---------------")

	//Created a go routine
	go func() {
		for i := 0; i < 3; i++ {
			//sending values in a loop 3 times
			fmt.Printf("sending %d\n", i)
			// to this channel
			ch <- i
			// then sleeping for 1 second and repeating the process
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 3; i++ {
		val := <-ch
		fmt.Printf("receieved %d\n", val)
	}
	elapsed := time.Since(start)

	fmt.Println(elapsed)
}
