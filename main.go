package main

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards.txt")
	cardsFromFile := newDeckFromFile("my_cards.txt")
	cardsFromFile.print()
}