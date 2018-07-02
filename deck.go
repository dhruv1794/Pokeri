package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Deck represents set of 52 cards
type Deck []Card

func (d *Deck) print() {
	for i, card := range *d {
		fmt.Println(i, card)
	}

}
func newDeck() Deck {
	cards := Deck{}
	card := Card{}

	cardSuites := []string{"Club", "Spade", "Diamond", "Heart"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, suite := range cardSuites {
		for _, value := range cardValues {

			card = Card{value: value, suite: suite}
			cards = append(cards, card)
		}
	}
	return cards
}

func (d Deck) shuffle() Deck {
	//We have to change seed value because by default it remains same
	//seed value is something that is the cause of randomeness in our program (rand.Intn uses this seed)
	//we need to make new rand type and make a new source
	//source is responsible for the seed value
	// random int 64 number helps create source (we will use time for this as it will always be unique (time.now))
	currentTime := time.Now()
	currentTimeINT64 := currentTime.UnixNano()
	source := rand.NewSource(currentTimeINT64)
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
	return d
}

//Deal takes in current deck and returns two decks one is one dealt card, and another is remaining cards
func (d *Deck) deal() (Deck, Card) {
	card := (*d)[0]

	*d = append((*d)[:0], (*d)[1:]...)

	return *d, card
}
