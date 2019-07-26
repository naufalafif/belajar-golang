package main

func main() {
	// var card string = "ace of spades"
	// card := newCard()

	// cards := newDect()
	// cards.saveToFile("my_cards")

	deckFromFile := newDectFromFile("my_cards")
	deckFromFile.suffle()
	deckFromFile.print()
	// cardInString := cards.toString()

	// cardInByte := []byte(cardInString)
	// fmt.Println(cardInByte)

}

func newCard() string {
	return "five of diamods"
}
