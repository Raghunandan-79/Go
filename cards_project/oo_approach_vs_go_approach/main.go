package main

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six oF Spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
