package tui

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Hit     key.Binding
	Stand   key.Binding
	Restart key.Binding
	Help    key.Binding
	Quit    key.Binding
}

var keys = keyMap{
	Hit: key.NewBinding(
		key.WithKeys("h"),
		key.WithHelp("h", "hit"),
	),
	Stand: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "stand"),
	),
	Restart: key.NewBinding(
		key.WithKeys(" "),
		key.WithHelp("Space", "restart"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit, k.Help}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Hit, k.Stand, k.Restart},
		{k.Help, k.Quit},
	}
}
