package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func divmod(a int, b int) (int, int) {
	return a / b, a % b
}

func main() {

	value := add(4, 5)
	fmt.Println(value)

	div, mod := divmod(10, 5)

	fmt.Println("div= %d , mod= %d\n", div, mod)
}
