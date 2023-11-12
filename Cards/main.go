package main

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
	
	_, remainingDeck := deal(cards, 4)
	remainingDeck.saveToFile("my_cards")
	//hand.print()
	//fmt.Println(remainingDeck.toString())
	
}