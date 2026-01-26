package test

import (
	"fmt"

	"github.com/errium/hitme/internal/game"
)

func TestDeck() {
	deck := game.NewDeck()

	fmt.Println("Full Deck Contents:")
	for i, card := range deck.Cards {
		fmt.Printf("%d: %s of %s (Value: %d)\n",
			i+1,
			card.Rank,
			card.Suit,
			card.GetValue())
	}
	fmt.Printf("\nTotal cards: %d\n", len(deck.Cards))
}

func TestDeckShuffle() {
	deck := game.NewDeck()
	game.Shuffle(deck)

	fmt.Println("Full Deck Contents:")
	for i, card := range deck.Cards {
		fmt.Printf("%d: %s of %s (Value: %d)\n",
			i+1,
			card.Rank,
			card.Suit,
			card.GetValue())
	}
	fmt.Printf("\nTotal cards: %d\n", len(deck.Cards))
}
