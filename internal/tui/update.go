package tui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/errium/hitme/internal/game"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Hit):
			if !m.game.Flow.IsGameOver() {
				m.game.Hit()
			}
		case key.Matches(msg, m.keys.Stand):
			if !m.game.Flow.IsGameOver() {
				m.game.Stand()
			}
		case key.Matches(msg, m.keys.Restart):
			if m.game.Flow.IsGameOver() {
				m.game = game.NewGame()
				m.game.Deal()
			}
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}
	}
	return m, nil
}
