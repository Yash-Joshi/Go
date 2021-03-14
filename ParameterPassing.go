package main

import (
	"fmt"
)

// Here the value slice passed is of reference type , because Go passes by referece slices.
// It also updates the i index value in the 0 based index
func doubleAt(value []int, i int) {
	value[i] *= i
}

//here in this function we are passing by value
func double(val int, n *int) (int, int) {
	p := val
	fmt.Println("Printing the value", p)

	*n *= p
	return p, *n
}

//in this function we are passing the value by reference
func doublePtr(n *int) {
	fmt.Println("Prior value", *n)
	*n *= 2
}

func main() {
	value := []int{1, 2, 3, 4, 5}
	doubleAt(value, 2)
	fmt.Println(value)

	// passing by value
	val := 10
	double(val, &val)
	fmt.Println(val, val)

	//passing by reference
	valP := 2
	doublePtr(&valP)
	fmt.Println(valP)

}
