package main

import "fmt"

// New deck type, slice of strings

type deck []string

func new_deck() deck {
	cards := deck{}

	card_suits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	card_value := []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}

	for i, suit := range card_suits {
		for j, value := range card_value {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}

}
