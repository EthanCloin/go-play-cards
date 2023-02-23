package main

import "fmt"

// a deck represents a slice of playing cards, represented
// simply as strings of the form "<suit> of <rank>"
type deck []string

/* 
receiver syntax: the '(d deck)' preceding func name
allows all `deck` type to utilize this method via dot operator
aka myDeck.print()
*/

// prints the card value and its zero-based index in deck
func (d deck) print(){
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// generates a new deck of 52 cards
func newDeck() deck {
	suits := []string{"Hearts", "Spades", "Clubs", "Diamonds"}
	ranks := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	newDeck := deck{}

	// go compiler refuses to allow any unutilized variables, 
	// assign any unneeded variables as '_' to allow compilation
	for _, suit := range suits {
		for _, rank := range ranks {
			newDeck = append(newDeck, rank + " of " + suit)
		}
	}
	return newDeck
}

// use slicing syntax to remove `handSize` cards from 
// deck. returns hand and updatedDeck
func dealHand(d deck, handSize int) (deck, deck){
	return d[:handSize], d[handSize:]
}