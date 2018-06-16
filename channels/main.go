package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, url := range urls {
		checkLink(url)
	}
}

func checkLink(url string) {
	_, err := http.Get(url) // blocking call
	if err != nil {
		fmt.Println(url, "might be down!")
	} else {
		fmt.Println(url, "is up!")
	}
}
