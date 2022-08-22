package UI

import (
	"time"

	"github.com/Grimmr/TheMagician/BackendInterface"
	"github.com/Grimmr/TheMagician/YGO"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type BackendSetupModel struct {
	spinner     spinner.Model
}

func NewBackendSetupModel() BackendSetupModel {
	//setup spinner
	spin := spinner.New()
	spin.Spinner = spinner.Moon
	spin.Spinner.FPS = time.Second / 8
	spin.Style = styleForegroundBright

	return BackendSetupModel{spinner: spin}
}

func (m BackendSetupModel) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, selectBackendCmd)
}

func (m BackendSetupModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case backendSelectedMsg:
		newModel := NewSetListSelectModel(msg.backend, m.spinner)
		return newModel, newModel.Init()
	}

	return m, nil
}

func (m BackendSetupModel) View() string {
	return m.spinner.View() + m.spinner.Style.Render(" Loading Set List Please Wait...")
}

type backendSelectedMsg struct {
	backend BackendInterface.Backend
}

func selectBackendCmd() tea.Msg {
	time.Sleep(3 * time.Second)
	return backendSelectedMsg{backend: YGO.YgoBackend{}}
}
