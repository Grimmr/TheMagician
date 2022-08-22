package main

import (
	"github.com/Grimmr/TheMagician/UI"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(UI.NewBackendSetupModel())
	p.Start()
}
