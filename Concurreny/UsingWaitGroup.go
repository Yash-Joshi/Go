//Get content type of sites by running the bunch of url

package main

import (
	"fmt"
	"net/http"
	"sync"
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
	//Creating a waitGroup
	var waitHere sync.WaitGroup
	for _, url := range urls {

		// So for every go routing we are going to create we are going to tell the waitGroup that
		// we have added another worker
		waitHere.Add(1)
		go func(url string) {
			returnType(url)
			waitHere.Done()
		}(url)

	}
	waitHere.Wait()
	// Above will output
	//https://golang.org -> text/html; charset=utf-8
	//https://api.github.com -> application/json; charset=utf-8
	//https://httpbin.org/xml -> application/xml
	elapsed := time.Since(start)

	fmt.Println(elapsed)
}
