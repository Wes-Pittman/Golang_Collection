package main

func main() {
	// var card string = "Aces of Spades"
	cards := deck{"Ace of Hearts", new_card()}
	cards = append(cards, "Six of Spades")

	cards.print()
}
func new_card() string {
	return "Five of Spades"
}
