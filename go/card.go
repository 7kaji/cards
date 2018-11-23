package main

import (
	"fmt"
	"strconv"
)

type card string

func newCard(suit string, number int) string {
	formatted_number := fmt.Sprintf("% 2s", strconv.Itoa(number))
	return formatted_number + " of " + suit
}
