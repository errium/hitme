package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/errium/hitme/internal/game"
)

type Model struct {
	game *game.Game
}

func New() Model {
	return Model{
		game: game.NewGame(),
	}
}

func (m Model) Init() tea.Cmd {
	m.game.Deal()
	return nil
}
