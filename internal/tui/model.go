package tui

import tea "github.com/charmbracelet/bubbletea"

type Model struct{}

func New() Model {
	return Model{}
}

func (m Model) Init() tea.Cmd {
	return nil
}
