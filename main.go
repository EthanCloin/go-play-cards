package main

import "fmt"

func main() {
	/* 
	interesting that this fxn declared in deck file doesn't need
	explicitly imported. 
	what happens if i have two fxns of same name in different files?
	*/
	cards := newDeck()

	// notice multiple values returned from single function <3
	myHand, cards := dealHand(cards, 5)
	myHand.print()

	fmt.Println("REMAINING IN DECK: ")
	cards.print()

	fmt.Println(cards.toString())
}