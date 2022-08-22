package UI

import (
	"fmt"

	"github.com/Grimmr/TheMagician/BackendInterface"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type modelState int

type setListSelectModel struct {
	backend           BackendInterface.Backend
	spinner           spinner.Model
	setListData       []map[string]string
	waitingForSetData bool
	setListUI         list
}

func NewSetListSelectModel(backEnd BackendInterface.Backend, spin spinner.Model) setListSelectModel {
	//setup setlist
	listUi := newList([]item{}, 20, styleForegroundBright, styleForeground)

	return setListSelectModel{backend: backEnd, spinner: spin, setListUI: listUi, waitingForSetData: true}
}

func (m setListSelectModel) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, m.fetchSetList())
}

func (m setListSelectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
		if !m.waitingForSetData {
			m.setListUI.Update(msg)
		}
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case fetchSetListDone:
		m.waitingForSetData = false
		m.setListData = msg.data
		var setItems []item
		for _, row := range m.setListData {
			setItems = append(setItems, item{name: row["name"], desc: row["date"]})
			m.setListUI.rows = setItems
		}
	}

	return m, nil
}

func (m setListSelectModel) View() string {
	if m.waitingForSetData {
		return m.spinner.View() + m.spinner.Style.Render(" Loading Set List Please Wait...")
	} else {
		return fmt.Sprintf(m.setListUI.View())
	}
}

type fetchSetListDone struct {
	data []map[string]string
}

func (m setListSelectModel) fetchSetList() tea.Cmd {
	return func() tea.Msg {
		return fetchSetListDone{data: m.backend.GetSetUrls(m.backend.GetSetListUrl())}
		/*return fetchSetListDone{data: []map[string]string{
			{"name": "pokemon", "date": "2022-March-01", "url": "dont care"},
			{"name": "magic", "date": "2012-June-08", "url": "dont care"},
			{"name": "digimon", "date": "2012-October-08", "url": "dont care"},
			{"name": "ygo", "date": "2032-October-18", "url": "dont care"},
		}}*/
	}
}
