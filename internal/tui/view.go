package tui

import (
	"fmt"

	"github.com/errium/hitme/internal/game"
)

func (m Model) View() string {
	s := "=== Blackjack ===\n\n"

	// Player hand
	s += fmt.Sprintf("Your hand: %v (Score: %d)\n", formatHand(m.game.PlayerHand), m.game.PlayerHand.Value())

	// Dealer hand
	if m.game.Flow.IsGameOver() {
		s += fmt.Sprintf("Dealer hand: %v (Score: %d)\n", formatHand(m.game.DealerHand), m.game.DealerHand.Value())
	} else {
		cards := m.game.DealerHand.Cards
		s += fmt.Sprintf("Dealer shows: %s %s [?]\n", cards[0].Rank, cards[0].Suit)
	}

	s += "\n"

	// Game result
	if m.game.Flow.IsGameOver() {
		switch m.game.Flow.Result() {

		case game.ResultPlayerWin:
			s += "🎉 YOU WIN!\n"

		case game.ResultDealerWin:
			s += "💀 DEALER WINS\n"

		case game.ResultTie:
			s += "🤝 PUSH (Tie)\n"
		}
		s += "\nPress q to quit"
	} else {
		s += "Press H to hit, S to stand, Q to quit"
	}

	return s
}

func formatHand(h *game.Hand) string {
	var cards []string

	for _, c := range h.Cards {
		cards = append(cards, fmt.Sprintf("%s %s", c.Rank, c.Suit))
	}

	return fmt.Sprintf("%v", cards)
}
