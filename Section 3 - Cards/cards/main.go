package main

//import "fmt"

//import "fmt"

func main() {
	// //var card string = "Ace of Spades"
	// //card := newCard()
		// ^Go will infer the type of a variable for us. Cool!
	// //fmt.Println(card)
	// //cards := []string{"Ace of Diamonds", newCard()}
	// cards := deck{"Ace of Diamonds", newCard()}

	// cards = append(cards, "Six of Spades")
		// ^The act of appending to an existing slice will return a NEW SLICE

		// This syntax is similar to for loops using enumerate() in python
	// // for i, card := range cards {
	// // 	fmt.Println(i, card)
	// // }
	// cards.print()

	// cards := newDeck()
	// //cards.print()
		// Syntax here is same as unpacking assignments in Python
	// //hand, remainingDeck := deal(cards, 5)

	// // hand.print()
	// // remainingDeck.print()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")

	// cards2 := newDeckFromFile("my_cards")
	// cards2.shuffle()
	// cards2.print()

	shufCards := newDeck()
	shufCards.shuffle()
	shufCards.print()
}
