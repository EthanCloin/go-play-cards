package main

import "fmt"

type deck []string

// receiver (allows all `deck` type to utilize this method)
func (d deck) print(){
	for i, card := range d {
		fmt.Println(i, card)
	}
}