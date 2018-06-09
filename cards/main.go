package main

// this required import was automatically added by VS code on save

/*
	Go is a statically typed language similar to C++ or Java

	Basic Go Types:
		- bool
		- string
		- int
		- float64
*/
func main() {
	// string type is inferred by the assigned value
	// var card = "Ace of Spades"
	//card := "Ace of Spades" // ONLY for a new variable
	//card = "Five of Diamonds" // just assign a new value
	//card := newCard()

	/*
		Arrays vs Slices
		Arrays:
			a static size
		Slices:
			can shrink or grow
			all items must be same datatype
	*/

	// define a string slice
	//cards := []string{"Ace of Diamonds", newCard()}
	//cards := deck{"Ace of Diamonds", newCard()}

	// does not modify slice, instead returns a new slice
	//cards = append(cards, "Six fo Spades")

	// iterate over slice and print each card
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// All cards
	cards := newDeckFromFile("my_cards")

	cards.print()

	cards.shuffle()

	cards.print()

	// return a hand and the rest of the cards
	//hand, remainingCards := deal(cards, 5)

	//hand.print()
	//remainingCards.print()

	//hand.saveToFile()
	cards.saveToFile("my_cards")
}

// a function that returns a string
// func newCard() string {
// 	return "Five of Diamonds"
// }
