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
		/*
			By default GO attempts to only use one core:
			go keyword will allow a different go routine to run
			when one blocks.  So still only one routine runs at a
			time when on a single core

			When GO is setup to use multiple cores
			each GO routine will get a core and can run
			in parallel

			CONCURRENCY IS NOT PARALLELISM
			(hyperthreading) vs (multiple cores)

			Main Routine exists after launching go routines
			and then exits without showing any output
		*/
		go checkLink(url)
	}
}

func checkLink(url string) {
	_, err := http.Get(url) // blocking call
	if err != nil {
		fmt.Println(url, "might be down!")
	}
	fmt.Println(url, "is up!")
}
