package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of "deck"
// which is a slice of strings
// - Can think of this statement as saying that this new deck type
//	 extends/borrows all of the behavior from a slice of string
type deck []string

// Constructor
func newDeck() deck {
	cards := deck{}
	
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, val := range cardValues {
			cards = append(cards, val+" of "+suit)
		}
	}

	return cards
}

// The (d deck) syntax indicates that we want this
// function to be a "receiver" for a deck.
func (d deck) print() {
	// . You can think of this "d" variable as analagous to "this" or "self"
	// . By convention, we usually name the calling type (in this case, a deck)
	//   using the first letter or first two letters of the type name.
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// This is like a static method?
/**
@return A deck representing the hand that was dealt and
		another deck representing the remaining deck
**/
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d)-1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}