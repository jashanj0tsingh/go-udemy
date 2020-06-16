package main

import "fmt"

// represents a deck of cards
type deck []string

// receiver functions
func (d deck) print()  {
	for _, card := range d {
		fmt.Println(fmt.Sprintf(`%s`, card))
	}
}

// returns a new deck of cards
func newDeck() deck {

	//
	cardSuits := [4]string{"\u2660", "\u2665", "\u2666", "\u2663"}
	cardRanks := [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	// create an empty deck
	cards := deck{}

	for _, suit := range cardSuits {
		for _, rank := range cardRanks {
			card := suit + " " + rank
			cards = append(cards, card)
		}
	}
	return cards
}