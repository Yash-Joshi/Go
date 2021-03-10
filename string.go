package main

import (
	"fmt"
)

func main() {
	book := "This is a new book"
	fmt.Println(book)

	//priting the length of the string
	fmt.Println(len(book))

	fmt.Println("book[0] = %v (type %T)\n", book[0], book[0])

	//uint8 is a byte

	//String in go is immutable
	//book[0] = 116 wont work

	//Slice (start, end) , ) based half empty range so we get 4 in the print index but we wont get 11
	fmt.Println(book[4:11])

	//Slice no end , here we get everything from charater number 4 to end
	fmt.Println(book[4:])

	//Slice no start everything upto 4 but not 4
	fmt.Println(book[:4])

	//Use + to concatenate strings
	fmt.Println("t" + book[1:])

	//Unicode
	fmt.Println("This is a ^ character")

	//Multi-line
	poem := `
	.. Far from the earth
	I am travelling to find you 
	Just to be with you 
	sing with you 
	and Love you
	oh Go, please do go
	`

	fmt.Println(poem)
}
