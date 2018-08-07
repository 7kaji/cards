package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardsSuits := []string{"♠", "♥", "♦", "♣"}
	cardsNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	for _, suit := range cardsSuits {
		for _, number := range cardsNumbers {
			card := newCard(suit, number)
			cards = append(cards, card)
		}
	}

	return cards
}

// d(cards) が self みたいな感じ
func (d deck) show() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), "\n")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
