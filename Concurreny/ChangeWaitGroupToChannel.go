//Get content type of sites by running the bunch of url

package main

import (
	"fmt"
	"net/http"
	"time"
)

func returnType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	fmt.Printf("%s -> %s\n", url, ctype)
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}
	start := time.Now()

	ch := make(chan string)

	go func() {
		for i := 0; i < len(urls); i++ {
			ch <- urls[i]
		}
	}()

	for i := 0; i < len(urls); i++ {
		val := <-ch

		fmt.Println(val)
	}
	// Above will output
	//https://golang.org -> text/html; charset=utf-8
	//https://api.github.com -> application/json; charset=utf-8
	//https://httpbin.org/xml -> application/xml
	elapsed := time.Since(start)

	fmt.Println(elapsed)
}
