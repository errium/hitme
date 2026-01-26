package game

import "strconv"

type Card struct {
	Suit string
	Rank string
}

func NewCard(Suit, Rank string) Card {
	return Card{Suit: Suit, Rank: Rank}
}

func (c Card) GetValue() int {
	switch c.Rank {
	case "a":
		return 11 // TODO: handle ace 1 & 11 logic
	case "k", "q", "j", "10":
		return 10
	default:
		v, _ := strconv.Atoi(c.Rank)
		return v
	}
}
