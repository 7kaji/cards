package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck lengh of 52, but got %v", len(d))
	}

	if d[0] != " 1 of ♠" {
		t.Errorf("Exptected first card of 1 of ♠, but got %v", d[0])
	}

	if d[len(d)-1] != "13 of ♣" {
		t.Errorf("Exptected last card of 13 of ♣, but got %v", d[len(d)-1])
	}
}

func TestShuffle(t *testing.T) {
	d := newDeck()
	d.shuffle()

	if len(d) != 52 {
		t.Errorf("Expected deck lengh of 52, but got %v", len(d))
	}

	if d[0] == " 1 of ♠" {
		t.Errorf("Exptected first card of 1 of ♠, but got %v", d[0])
	}

	if d[len(d)-1] == "13 of ♣" {
		t.Errorf("Exptected last card of 13 of ♣, but got %v", d[len(d)-1])
	}
}
