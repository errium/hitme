package tui

import (
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/errium/hitme/internal/game"
)

type Model struct {
	game   *game.Game
	keys   keyMap
	help   help.Model
	width  int
	height int
}

func NewModel() Model {
	return Model{
		game: game.NewGame(),
		keys: keys,
		help: help.New(),
	}
}

func (m Model) Init() tea.Cmd {
	m.game.Deal()
	return nil
}
