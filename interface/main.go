package main

import (
	"fmt"
)

/*
	Notes on Interfaces:

	type bot interface {
	//   ^interface name
		getGreeting			(string, int)	(string, error)
	//	  ^function name     ^list of args 	^list of return types
	}
*/

type bot interface {
	getGreeting() string
}

// englishBot and spanishBot got the bot interace implicitly by
// the usage of the print function
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi There!"
}

// when not actually using the receiver, can omit
// the variable and just leave the type instead
func (spanishBot) getGreeting() string {
	return "Hola!"
}
