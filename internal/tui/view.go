package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/errium/hitme/internal/game"
)

var (
	headerStyle = lipgloss.NewStyle().
			Padding(0, 1, 0, 1).
			Bold(true).
			Foreground(lipgloss.ANSIColor(15)).
			Background(lipgloss.ANSIColor(5))
	commonTextStyle = lipgloss.NewStyle().
			Bold(true)
	scoreStyle = lipgloss.NewStyle().
			Faint(true)
	resultStyle = lipgloss.NewStyle().
			Padding(0, 1, 0, 1).
			Bold(true).
			Foreground(lipgloss.ANSIColor(5)).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.ANSIColor(8))
)

const header string = `hitme`

func (m Model) View() string {
	top := viewTop()
	center := viewCenter(m)
	bottom := viewBottom(m)

	centeredMiddle := lipgloss.Place(
		m.width,
		m.height-lipgloss.Height(top)-lipgloss.Height(bottom),
		lipgloss.Center,
		lipgloss.Center,
		center,
	)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		top,
		centeredMiddle,
		bottom,
	)
}

func viewTop() string {
	s := headerStyle.Render(header) + "\n"
	return s
}

func viewCenter(m Model) string {
	var s string
	s += renderDealerHand(m)
	s += "\n\n\n"
	s += renderPlayerHand(m)
	s += "\n\n"
	s += renderGameResult(m)
	return s
}

func viewBottom(m Model) string {
	s := m.help.View(m.keys)
	return s
}

// NOTE: Below this comment are helper functions

func renderPlayerHand(m Model) string {
	s := commonTextStyle.Render("Your hand:")
	s += "\n"
	s += renderHand(m.game.PlayerHand.Cards, false)
	s += "\n"

	score := fmt.Sprintf("Score: %d", m.game.PlayerHand.Value())
	s += scoreStyle.Render(score)
	return s
}

func renderDealerHand(m Model) string {
	s := commonTextStyle.Render("Dealer's hand:")
	s += "\n"
	hideSecond := !m.game.Flow.IsGameOver()
	s += renderHand(m.game.DealerHand.Cards, hideSecond)
	s += "\n"

	if m.game.Flow.IsGameOver() {
		score := fmt.Sprintf("Score: %d", m.game.DealerHand.Value())
		s += scoreStyle.Render(score)
	} else {
		partialScore := m.game.DealerHand.Cards[0].Value()
		score := fmt.Sprintf("Score: ~%d", partialScore)
		s += scoreStyle.Render(score)
	}
	return s
}

func renderGameResult(m Model) string {
	if !m.game.Flow.IsGameOver() {
		return ""
	}

	var s string
	switch m.game.Flow.Result() {
	case game.ResultPlayerWin:
		s = "You won!"
	case game.ResultDealerWin:
		s = "Dealer won"
	case game.ResultPush:
		s = "Push (Tie)"
	default:
		s = ""
	}
	return resultStyle.Render(s)
}
