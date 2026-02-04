package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/errium/hitme/internal/game"
)

const (
	cardWidth    int    = 8
	cardHeight   int    = 5
	hiddenSymbol string = "?"
)

var (
	generic = lipgloss.NewStyle().
		Align(lipgloss.Top, lipgloss.Left).
		Border(lipgloss.RoundedBorder()).
		Bold(true).
		Padding(0, 1).
		Width(cardWidth).
		Height(cardHeight)
	blackCardStyle = generic.
			Foreground(lipgloss.ANSIColor(7))
	redCardStyle = generic.
			Foreground(lipgloss.ANSIColor(1))
	hiddenCardStyle = generic.
			Foreground(lipgloss.ANSIColor(5))
)

func getSuitSymbol(suit string) string {
	switch suit {
	case "clubs":
		return "♧"
	case "diamonds":
		return "♢"
	case "hearts":
		return "♡"
	case "spades":
		return "♤"
	default:
		return "?"
	}
}

func renderCard(c game.Card) string {
	rank := strings.ToUpper(c.Rank)
	suit := getSuitSymbol(c.Suit)
	horizontalSpace := cardWidth - len(rank) - len(suit)
	verticalSpace := cardHeight - 1

	var style lipgloss.Style
	if c.Suit == "hearts" || c.Suit == "diamonds" {
		style = redCardStyle
	} else {
		style = blackCardStyle
	}

	content := rank + strings.Repeat(" ", horizontalSpace) + suit
	content += strings.Repeat("\n", verticalSpace)
	content += suit + strings.Repeat(" ", horizontalSpace) + rank
	return style.Render(content)
}

func renderHiddenCard() string {
	horizontalSpace := cardWidth - 4
	verticalSpace := cardHeight - 1

	content := hiddenSymbol + strings.Repeat(" ", horizontalSpace) + hiddenSymbol
	content += strings.Repeat("\n", verticalSpace)
	content += hiddenSymbol + strings.Repeat(" ", horizontalSpace) + hiddenSymbol
	return hiddenCardStyle.Render(content)
}

func renderHand(cards []game.Card, hideSecond bool) string {
	if len(cards) == 0 {
		return ""
	}

	var renderedCards []string
	for i, card := range cards {
		if i == 1 && hideSecond {
			renderedCards = append(renderedCards, renderHiddenCard())
		} else {
			renderedCards = append(renderedCards, renderCard(card))
		}
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, renderedCards...)
}
