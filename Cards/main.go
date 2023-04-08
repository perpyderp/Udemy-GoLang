package main

func main() {

	// var card string = "Ace of Spades"
	cards := deck{newCard(), newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
