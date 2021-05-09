//The build in select functio allows us to work with several channels together
//Everytime a channel with selecting on become free, either for sending or receieving,
// you ll do an action on this channel

package main

import (
	"fmt"
)

func main() {
	channelOne, channelTwo := make(chan int), make(chan int)

	go func() {
		channelOne <- 42
	}()

	select {
	case val := <-channelOne:
		fmt.Printf("got  %d from channel One \n", val)
	case val := <-channelTwo:
		fmt.Printf("got  %d from channel Two \n", val)
	}
}

//One of the common use of select is to impliment timeouts
