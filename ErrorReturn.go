//Error return from code

package main

import (
	"fmt"
	"math"
)

func sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0.0, fmt.Errorf("sqrt of negative value", a)
	}

	return math.Sqrt(a), nil

}

func main() {

	s1, err := sqrt(2.0)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(s1)
	}

	s2, err := sqrt(-2.0)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(s2)
	}

}
