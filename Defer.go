//Go has a garbage collector

// Defer run in reverse order
// It can be used to clean up resources, files, sockets, virtaul machines
package main

import (
	"fmt"
)

func cleanup(name string) {
	fmt.Println("Cleaning up", name)
}

func worker() {
	defer cleanup("A")

	fmt.Println("Worker")

	defer cleanup("A1")
}

func main() {
	worker()
}
