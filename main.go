package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	card := newCard()
	fmt.Println(card)

	cards := []string {newCard(), "Ace of Clubs"}
	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}