package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/errium/hitme/internal/game"
)

type Model struct {
	game     *game.Game
	quit     bool
	showHelp bool
}

func NewModel() *Model {
	return &Model{
		game:     game.NewGame(),
		quit:     false,
		showHelp: true,
	}
}

func (m Model) Init() tea.Cmd {
	m.game.Deal()
	return nil
}
