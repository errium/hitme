package main

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/errium/hitme/internal/tui"
)

func main() {
	p := tea.NewProgram(tui.New())
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
