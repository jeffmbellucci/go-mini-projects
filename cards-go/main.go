package main

func main() {
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()
	cards.print()
}

