package UI

import (
	"fmt"
	"time"

	"github.com/Grimmr/TheMagician/BackendInterface"
	"github.com/Grimmr/TheMagician/YGO"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type modelState int

const (
	modelState_LoadingSets = iota
	modelState_selectSet
)

type model struct {
	backend     BackendInterface.Backend
	state       modelState
	spinner     spinner.Model
	setListData []map[string]string
	setListUI   list
}

func NewModel() model {
	//setup spinner
	spin := spinner.New()
	spin.Spinner = spinner.Moon
	spin.Spinner.FPS = time.Second / 8
	spin.Style = styleForegroundBright

	//setup setlist
	listUi := newList([]item{}, styleForegroundBright, styleForeground)

	return model{backend: YGO.YgoBackend{}, state: modelState_LoadingSets, spinner: spin, setListUI: listUi}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, m.fetchSetList())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
		if m.state == modelState_selectSet {
			m.setListUI.Update(msg)
		}
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case fetchSetListDone:
		m.state = modelState_selectSet
		m.setListData = msg.data
		var setItems []item
		for _, row := range m.setListData {
			setItems = append(setItems, item{name: row["name"], desc: row["date"]})
			m.setListUI.rows = setItems
		}
	}

	return m, nil
}

func (m model) View() string {
	switch m.state {
	case modelState_LoadingSets:
		return m.spinner.View() + m.spinner.Style.Render(" Loading Set List Please Wait...")
	case modelState_selectSet:
		return fmt.Sprintf(m.setListUI.View())
	}
	return ""
}

type fetchSetListDone struct {
	data []map[string]string
}

func (m model) fetchSetList() tea.Cmd {
	return func() tea.Msg {
		//return fetchSetListDone{data: m.backend.GetSetUrls(m.backend.GetSetListUrl())}
		return fetchSetListDone{data: []map[string]string{
			{"name": "pokemon", "date": "2022-March-01", "url": "dont care"},
			{"name": "magic", "date": "2012-June-08", "url": "dont care"}},
		}
	}
}
