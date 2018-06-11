package main

import (
	"fmt"
)

type englishBot struct{}

type spanishBot struct{}

func main() {
	eb := englishBot{}
	// sb := spanishBot{}

	printGreeting(eb)
	// printGreeting(sb)
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func (englishBot) getGreeting() string {
	return "Hi There!"
}

// when not actually using the receiver, can omit
// the variable and just leave the type instead
func (spanishBot) getGreeting() string {
	return "Hola!"
}
