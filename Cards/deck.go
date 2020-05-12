package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	// declare empty deck
	cards := deck{}

	// create colors and values
	cardSuites := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// iterate through both arrays to create every single card of a 52 card deck
	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Print function to print out all cards
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
