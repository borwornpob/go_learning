package main

import "fmt"

//var someValue int = 64

func main() {
	//var card string = "Ace of Spades" 
	//card := "Ace of Spades"
	//card := newCard()
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	cards := newDeck()
	
	hand, remainingDeck := deal(cards, 4)
	hand.print()
	fmt.Println()
	remainingDeck.print()
}