package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/errium/hitme/internal/game"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "h":
			if m.game.Flow.State() == game.StatePlayerTurn {
				m.game.Hit()
			}
		case "s":
			if m.game.Flow.State() == game.StatePlayerTurn {
				m.game.Stand()
			}
		case "r":
			if m.game.Flow.IsGameOver() {
				m.game = game.NewGame()
				m.game.Deal()
			}
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}
