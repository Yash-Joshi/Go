//Challenge is to find the maximum number present in the silce.
// 16,8,42,4,23,15

package main

import (
	"fmt"
)

func main() {

	nums := []int{16, 8, 42, 4, 23, 15}

	// my code here
	max := nums[0] // Initilize the first number to max
	//[1:] skips the first value
	for _, value := range nums[1:] {
		if value > max {
			max = value
		}

	}

	fmt.Println(max)

}
