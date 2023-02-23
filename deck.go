package main

import "fmt"

type deck []string

// receiver syntax: (d deck) preceding func name
// allows all `deck` type to utilize this method
func (d deck) print(){
	for _, card := range d {
		fmt.Println(card)
	}
}

func newDeck() deck {
	suits := []string{"Hearts", "Spades", "Clubs", "Diamonds"}
	ranks := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	newDeck := deck{}

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