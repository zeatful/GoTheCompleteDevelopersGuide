package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	/*
		Create a channel to use for communication between go routines and
		main routine

		Sending data with Channels:
		chanel <- 5
			send the value 5 into this channel

		myNumber <- channel
			wait for a value to be sent into the hannel.  When we get one,
			assign the value to myNumber

		fmt.Println(<- channel)
			wait for a value to be sent into the channel.  When we get one, log
			it out immediately
	*/
	c := make(chan string)

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

			Main Routine launches go routines
			and then exits without showing any output
			because it exits before the go routines finish
		*/
		go checkLink(url, c)
	}

	/*
		waiting for a channel message is a blocking call,
		after the first message is sent to the channel,
		the main routine prints it and exits before any other
		go routines finish...

		//receive a value from the channel and immediately log it
		fmt.Println(<-c)
	*/

	/*
		wait for channel to return a value, assign it to l
		then pass it to / spawn a new go routine
	*/
	for u := range c {
		//go checkLink(l, c)
		// function literal
		go func(url string) { // literal must take a parameter
			time.Sleep(5 * time.Second)
			checkLink(url, c)
		}(u) // passing the current u variable by value
		/* if you directly reference the l variable in the function literal,
		it may have been changed before it's used (IE: main function changes
		it in memory before go routine tries to use it)*/
	}
}

func checkLink(url string, c chan string) {
	_, err := http.Get(url) // blocking call
	if err != nil {
		// send message to channel
		fmt.Println(url, "might be down!")
		c <- url
	}
	// send message to channel
	fmt.Println(url, "is up!")
	c <- url
}
