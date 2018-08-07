package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.show()
	cards.saveToFile("deck.txt")
}
