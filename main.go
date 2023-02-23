package main

import "fmt"
func main() {
	// interesting that this fxn declared in deck file doesn't need
	// explicitly imported. what happens if i have two fxns of same 
	// name ?
	cards := newDeck()
	// cards.print()

	myHand, cards := dealHand(cards, 5)
	myHand.print()
	fmt.Println("REMAINING: ")
	cards.print()
}