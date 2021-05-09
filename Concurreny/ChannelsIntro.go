//Preferred way to communicate between go routines is by using channels.
// Channels are typed one direcrtional pipes
//You can send values at one end and receieve it from another end
//If you try to receive and there's nothing in the channel, you will get blocked.
//Kinds of channels : 2
// 1. Buffered 2. Unbuffered

package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	<-ch
	fmt.Println("Here")
}

// The output of this will give us
// fatal error : all goroutines are asleep -deadlock
// since, nothing was pushed to the channel
