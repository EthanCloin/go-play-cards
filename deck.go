package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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
	ranks := []string {
		"Ace", "Two", "Three", "Four", "Five", "Six", 
		"Seven", "Eight", "Nine", "Ten", "Jack",
		"Queen", "King",
	}

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

func (d deck) toString() string {
	return strings.Join(d, ",")
// 	stringified := ""

// 	for _, card := range d {
// 		stringified += card
// 	}

// 	return stringified
}

func (d deck) saveToFile(filename string) error{
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() {
	for i := range d {
		// generate seed based on curtime
		src := rand.NewSource(time.Now().UnixMicro())
		r := rand.New(src)
		// generate index based on seed
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}