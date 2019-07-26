package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of dek
// Which is a slice of strings

type dek []string

func (d dek) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d dek) toString() string {
	return strings.Join(d, ",")
}

func (d dek) saveToFile(filename string) error {
	error := ioutil.WriteFile(filename, []byte(d.toString()), 0666)
	return error
}

func deal(d dek, handSize int) (dek, dek) {
	return d[:handSize], d[handSize:]
}

func newDect() dek {
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Club"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	cards := dek{}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+` `+suit)
		}
	}

	return cards
}

func newDectFromFile(filename string) dek {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	oneStringDeck := string(bs)
	dekFromFile := strings.Split(oneStringDeck, ",")
	cards := dek(dekFromFile)
	return cards
}

// Notes, to make rando generate new random number
// we need to make sure generator always use new seed/or source of random generator
// random generator work in 3 phase, create seed, create source base on seed, create generator base on source

func (d dek) suffle() dek {
	seed := time.Now().UnixNano()  // Seed
	source := rand.NewSource(seed) // Create Random Source
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
	return d
}
