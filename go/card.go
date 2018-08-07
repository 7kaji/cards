package main

import (
	"fmt"
	"strconv"
)

type card string

func newCard(suit string, number int) string {
	return fmt.Sprintf("% 2s", strconv.Itoa(number)+" of "+suit)
}
