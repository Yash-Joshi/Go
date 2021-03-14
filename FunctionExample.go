//Write a function that gets a URL and returns the value of Content-Type response HTTP header

//The function should return an error if it cant perform a GET request

// func contentType(url string) (string, error)
package main

import (
	"fmt"
	"net/http"
)

func contentType(t string) (string, error) {
	resp, err := http.Get(t)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	ctype := resp.Header.Get("Content-Type")
	if ctype == "" {
		return "", fmt.Errorf("cannot find Content-Type header")
	}

	return ctype, nil
}

func main() {
	anotherVar, err := contentType("https://golang.org")
	if err != nil {
		fmt.Println("error is returned -->", err)
	} else {

		fmt.Println(anotherVar)

	}
}
