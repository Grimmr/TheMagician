package UI

import (
	"fmt"
	"os"

	"github.com/Grimmr/TheMagician/BackendInterface"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type OutputCsvModel struct {
	backend  BackendInterface.Backend
	spinner  spinner.Model
	setUri   string
	outFile  string
	cardUrls []string
	cardData []map[string]string
	state    int
	progress int
}

func NewOutputCsvModel(backEnd BackendInterface.Backend, spin spinner.Model, uri string, file string) OutputCsvModel {
	//setup setlist
	return OutputCsvModel{backend: backEnd, spinner: spin, setUri: uri, state: 0, progress: 0, outFile: file}
}

func (m OutputCsvModel) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, m.fetchCardUrls())
}

func (m OutputCsvModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

	case fetchCardUrlsDone:
		m.cardUrls = msg.data
		m.state = 1
		return m, m.fetchData(m.cardUrls[0])

	case fetchCardDataDone:
		m.cardData = append(m.cardData, msg.data)
		m.progress++
		if m.progress >= len(m.cardUrls) {
			m.state = 2
			return m, m.outputCsv()
		} else {
			return m, m.fetchData(m.cardUrls[m.progress])
		}
	}

	return m, nil
}

func (m OutputCsvModel) View() string {
	if m.state == 0 {
		return m.spinner.View() + m.spinner.Style.Render(" Loading Card Data Please Wait...")
	} else if m.state == 1 {
		return m.spinner.View() + m.spinner.Style.Render(fmt.Sprintf(" %d/%d Loading Card Data Please Wait...", m.progress+1, len(m.cardUrls)))
	} else if m.state == 2 {
		return "dumping to " + m.outFile
	}
	return ""
}

type fetchCardUrlsDone struct {
	data []string
}

func (m OutputCsvModel) fetchCardUrls() tea.Cmd {
	return func() tea.Msg {
		return fetchCardUrlsDone{data: m.backend.GetCardUrlsFromSet(m.setUri)}
	}
}

type fetchCardDataDone struct {
	data map[string]string
}

func (m OutputCsvModel) fetchData(url string) tea.Cmd {
	return func() tea.Msg {
		return fetchCardDataDone{data: m.backend.GetCardDataFromUrl(url)}
	}
}

func (m OutputCsvModel) outputCsv() tea.Cmd {
	return func() tea.Msg {
		outstring := "cardname,cardid"
		for _, row := range m.cardData {
			outstring += fmt.Sprintf("\n%s,%s", row["name"], row["password"])
		}
		os.WriteFile(m.outFile, []byte(outstring), 0644)
		return tea.Quit
	}
}
