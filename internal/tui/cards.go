package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/errium/hitme/internal/game"
)

const (
	cardWidth  int = 8
	cardHeight int = 5
)

var (
	generic = lipgloss.NewStyle().
		Align(lipgloss.Top, lipgloss.Left).
		Bold(true).
		Border(lipgloss.RoundedBorder()).
		Padding(0, 1).
		Width(cardWidth).
		Height(cardHeight)
	blackCardStyle = generic.
			Foreground(lipgloss.ANSIColor(7))
	redCardStyle = generic.
			Foreground(lipgloss.ANSIColor(1))
	hiddenCardStyle = generic.
			Foreground(lipgloss.ANSIColor(3))
)

func renderCard(c game.Card) string {
	rank := strings.ToUpper(c.Rank)
	suit := getSuitSymbol(c.Suit)
	vSpaceBetween := cardWidth - len(rank) - len(suit)
	hSpaceBetween := cardHeight - 1

	var style lipgloss.Style
	if c.Suit == "hearts" || c.Suit == "diamonds" {
		style = redCardStyle
	} else {
		style = blackCardStyle
	}

	content := rank + strings.Repeat(" ", vSpaceBetween) + suit
	content += strings.Repeat("\n", hSpaceBetween)
	content += suit + strings.Repeat(" ", vSpaceBetween) + rank
	return style.Render(content)
}

func renderHiddenCard() string {
	vSpaceBetween := cardWidth - 4
	hSpaceBetween := cardHeight - 1

	content := "?" + strings.Repeat(" ", vSpaceBetween) + "?"
	content += strings.Repeat("\n", hSpaceBetween)
	content += "?" + strings.Repeat(" ", vSpaceBetween) + "?"
	return hiddenCardStyle.Render(content)
}

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

	return joinCardsHorizontally(renderedCards)
}

func joinCardsHorizontally(cards []string) string {
	if len(cards) == 0 {
		return ""
	}

	// Разбиваем каждую карту на строки
	cardLines := make([][]string, len(cards))
	for i, card := range cards {
		cardLines[i] = strings.Split(card, "\n")
	}

	// Склеиваем построчно
	var result []string
	height := len(cardLines[0])

	for row := 0; row < height; row++ {
		var line string
		for i, card := range cardLines {
			line += card[row]
			if i < len(cardLines)-1 {
				line += " " // пробел между картами
			}
		}
		result = append(result, line)
	}

	return strings.Join(result, "\n")
}
