package tui

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Hit     key.Binding
	Stand   key.Binding
	Restart key.Binding
	Help    key.Binding
	Quit    key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Hit, k.Stand},
		{k.Help, k.Restart, k.Quit},
	}
}

var keys = keyMap{
	Hit: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "hit"),
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
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}
