package game

import "strconv"

type Card struct {
	Suit string
	Rank string
}

const (
	SuitClubs    = "clubs"
	SuitDiamonds = "diamonds"
	SuitHearts   = "hearts"
	SuitSpades   = "spades"
)

const (
	RankAce   = "a"
	RankKing  = "k"
	RankQueen = "q"
	RankJack  = "j"
	Rank10    = "10"
	Rank9     = "9"
	Rank8     = "8"
	Rank7     = "7"
	Rank6     = "6"
	Rank5     = "5"
	Rank4     = "4"
	Rank3     = "3"
	Rank2     = "2"
)

func NewCard(suit, rank string) Card {
	return Card{Suit: suit, Rank: rank}
}

func (c Card) Value() int {
	switch c.Rank {
	case RankAce:
		return 11

	case RankKing, RankQueen, RankJack, Rank10:
		return 10

	default:
		value, _ := strconv.Atoi(c.Rank)
		return value
	}
}
