package tui

import (
	"fmt"

	"github.com/errium/hitme/internal/game"
)

const header string = `╻ ╻╻╺┳╸┏┳┓┏━╸
┣━┫┃ ┃ ┃┃┃┣╸ 
╹ ╹╹ ╹ ╹ ╹┗━╸`

func (m Model) View() string {
	s := header + "\n\n"

	if m.game.Flow.IsGameOver() {
		s += fmt.Sprintf("Dealer hand: %v (Score: %d)\n", formatHand(m.game.DealerHand), m.game.DealerHand.Value())
	} else {
		cards := m.game.DealerHand.Cards
		s += fmt.Sprintf("Dealer shows: %s %s [?]\n", cards[0].Rank, cards[0].Suit)
	}

	s += fmt.Sprintf("Your hand: %v (Score: %d)\n", formatHand(m.game.PlayerHand), m.game.PlayerHand.Value())

	s += "\n"

	if m.game.Flow.IsGameOver() {
		switch m.game.Flow.Result() {
		case game.ResultPlayerWin:
			s += "You won!\n"
		case game.ResultDealerWin:
			s += "Dealer won\n"
		case game.ResultPush:
			s += "Push (Tie)\n"
		}
	}

	s += "\n" + m.help.View(m.keys)

	return s
}

func formatHand(h *game.Hand) string {
	var cards []string

	for _, c := range h.Cards {
		cards = append(cards, fmt.Sprintf("%s %s", c.Rank, c.Suit))
	}

	return fmt.Sprintf("%v", cards)
}
