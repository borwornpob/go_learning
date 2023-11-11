package main

import "fmt"

// Create a new type of 'deck'
// Which is a slice of stings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValue := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "Kings"}

	for _, suit := range cardSuits {
		for _, value := range cardValue {
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}

// Print every cards inside deck
func (d deck) print() {
	for i, card := range d{
		fmt.Println(i, card)
	}
}

// standalone function
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

