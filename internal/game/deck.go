package game

import (
	"math/rand/v2"
)

type Deck struct {
	cards []Card
}

func NewDeck() *Deck {
	suits := []string{
		SuitClubs, SuitDiamonds,
		SuitHearts, SuitSpades,
	}
	ranks := []string{
		RankAce, RankKing, RankQueen, RankJack,
		Rank10, Rank9, Rank8, Rank7, Rank6,
		Rank5, Rank4, Rank3, Rank2,
	}
	d := &Deck{}

	for _, suit := range suits {
		for _, rank := range ranks {
			d.cards = append(d.cards, NewCard(suit, rank))
		}
	}
	return d
}

func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

func (d *Deck) Draw() Card {
	if len(d.cards) == 0 {
		panic("Deck was empty when .Draw() was called.")
	}
	card := d.cards[0]
	d.cards = d.cards[1:]
	return card
}
