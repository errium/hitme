package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/errium/hitme/internal/game"
)

var headerStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.ANSIColor(15)).
	Background(lipgloss.ANSIColor(9))

const header string = `╻ ╻╻╺┳╸┏┳┓┏━╸
┣━┫┃ ┃ ┃┃┃┣╸ 
╹ ╹╹ ╹ ╹ ╹┗━╸`

func (m Model) View() string {
	s := headerStyle.Render(header) + "\n\n"

	// Player hand
	s += "Your hand:\n"
	s += renderHand(m.game.PlayerHand.Cards, false)
	s += fmt.Sprintf("\nScore: %d\n\n", m.game.PlayerHand.Value())

	// Dealer hand
	s += "Dealer hand:\n"
	hideSecond := !m.game.Flow.IsGameOver()
	s += renderHand(m.game.DealerHand.Cards, hideSecond)

	if m.game.Flow.IsGameOver() {
		s += fmt.Sprintf("\nScore: %d\n\n", m.game.DealerHand.Value())
	} else {
		s += "\n\n"
	}

	s += winDisplay(m)
	s += "\n" + m.help.View(m.keys)

	return s
}

func winDisplay(m Model) string {
	var s string
	if m.game.Flow.IsGameOver() {
		switch m.game.Flow.Result() {
		case game.ResultPlayerWin:
			s += "🎉 You won!\n"
		case game.ResultDealerWin:
			s += "💀 Dealer won\n"
		case game.ResultPush:
			s += "🤝 Push (Tie)\n"
		}
	}
	return s
}
