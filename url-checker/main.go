package main

import (
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"https://www.",
		"https://www.",
		"https://www.",
		"https://www.",
		"https://www.",
	}

	c := make(chan string)
	e := make(chan string)

	for _, url := range urls {
		time.Sleep(5 * time.Minute)
		go checkUrl(url, c, e)
	}

	for u := range c {
		go func(url string) {
			go checkUrl(url, c, e)
		}(u)
	}
}

func checkUrl(url string, c chan string, e chan string) {
	_, err := http.Get(url)
	if err != nil {
		c <- url
	} else {
		e <- url
	}
}
