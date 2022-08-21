package UI

import (
	"math"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type item struct {
	name string
	desc string
}

type list struct {
	pos             int
	rows            []item
	selectedStyle   lipgloss.Style
	unselectedStyle lipgloss.Style
}

func newList(items []item, s lipgloss.Style, u lipgloss.Style) list {
	return list{pos: 0, rows: items, selectedStyle: s, unselectedStyle: u}
}

func (this *list) Update(msg tea.Msg) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			this.pos -= 1
			if this.pos < 0 {
				this.pos = 0
			}
		case "down":
			this.pos += 1
			if this.pos >= len(this.rows)-1 {
				this.pos = len(this.rows) - 1
			}
		}
	}
}

func (this *list) View() string {
	var body string
	for index, row := range this.rows {
		seleceted := index == this.pos
		if seleceted {
			body += this.selectedStyle.Render(" > "+row.name+" ("+row.desc+")") + "\n"
		} else {
			body += this.unselectedStyle.Render("   "+row.name+" ("+row.desc+")") + "\n"
		}
	}

	//find width of table body
	tail := "Enter: Select Set   Left: Prev Page   Right: Next Page"
	paddingAmount := lipgloss.Width(body) - lipgloss.Width(tail)
	if paddingAmount > 0 {
		lPad := math.Ceil(float64(paddingAmount) / 2.0)
		rpad := math.Floor(float64(paddingAmount) / 2.0)
		tail = this.unselectedStyle.Render(strings.Repeat("-", int(lPad)) + tail + strings.Repeat("-", int(rpad)))
	} else {
		tail = this.unselectedStyle.Render(tail)
	}

	return body + tail
}
