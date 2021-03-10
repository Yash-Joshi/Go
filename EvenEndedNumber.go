//Challenge
// A even ended number is a number  with the same first digit  and last digits
// example 1,11,121 etc

// How many even-ended numbers results from  multiplying  two four digits numbers ?
// example  1001 X 1011  = 1012011  even ended number

// In short check if a number is pallindrom or not !

// To check --> convert the number to string  --> fmt.Sprintf()

package main

import (
	"fmt"
)

func main() {

	n := 42
	// We use Sprintf to convert n to string
	s := fmt.Sprintf("%d", n)

	// We print value and type of s
	fmt.Printf("s = %s (type %T)\n ", s, s)

	//verb q -- add quotes in the string
	fmt.Printf("s = %q (type %T)\n", s, s)

}
