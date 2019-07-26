package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of deck
// Which is a slice of strings

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	error := ioutil.WriteFile(filename, []byte(d.toString()), 0666)
	return error
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func newDect() deck {
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Club"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	cards := deck{}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+` `+suit)
		}
	}

	return cards
}

func newDectFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	oneStringDeck := string(bs)
	deckFromFile := strings.Split(oneStringDeck, ",")
	cards := deck(deckFromFile)
	return cards
}

// Notes, to make rando generate new random number
// we need to make sure generator always use new seed/or source of random generator
// random generator work in 3 phase, create seed, create source base on seed, create generator base on source

func (d deck) suffle() deck {
	seed := time.Now().UnixNano()  // Seed
	source := rand.NewSource(seed) // Create Random Source
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
	return d
}
