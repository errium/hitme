package main

import (
	"log"
	"os"

	"github.com/errium/hitme/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(tui.NewModel())
	if _, err := p.Run(); err != nil {
		log.Fatal()
		os.Exit(1)
	}
}
