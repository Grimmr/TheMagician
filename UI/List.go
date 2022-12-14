package UI

import (
	"fmt"
	"math"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type item struct {
	name string
	desc string
}

type list struct {
	page            int
	pageSize        int
	pos             int
	rows            []item
	selectedStyle   lipgloss.Style
	unselectedStyle lipgloss.Style
	headerStyle     lipgloss.Style
}

func newList(items []item, pageS int, s lipgloss.Style, u lipgloss.Style, h lipgloss.Style) list {
	return list{pos: 0, rows: items, pageSize: pageS, selectedStyle: s, unselectedStyle: u, headerStyle: h}
}

func (this *list) pageCount() int {
	return int(math.Ceil(float64(len(this.rows)) / float64(this.pageSize)))
}

func (this *list) currentPageSize() int {
	if this.page < this.pageCount()-1 {
		return this.pageSize
	} else {
		return len(this.rows) % this.pageSize
	}
}

func (this *list) PageUp() {
	if this.page > 0 {
		this.page -= 1
	} else {
		this.page = this.pageCount() - 1
	}
}

func (this *list) PageDown() {
	if this.page < this.pageCount()-1 {
		this.page += 1
	} else {
		this.page = 0
	}
}

func (this *list) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			this.pos -= 1
			if this.pos < 0 {
				this.PageUp()
				this.pos = this.currentPageSize() - 1
			}
		case "down":
			this.pos += 1
			if this.pos >= this.currentPageSize() {
				this.PageDown()
				this.pos = 0
			}
		case "left":
			this.PageUp()
			if this.pos >= this.currentPageSize() {
				this.pos = this.currentPageSize() - 1
			}
		case "right":
			this.PageDown()
			if this.pos >= this.currentPageSize() {
				this.pos = this.currentPageSize() - 1
			}
		case "enter":
			return this.SelectItem(this.page*this.pageSize + this.pos)
		}
	}

	return nil
}

type selectedItemFromList struct {
	listUI *list
	choice item
	index  int
}

func (this *list) SelectItem(choice int) tea.Cmd {
	return func() tea.Msg {
		return selectedItemFromList{listUI: this, choice: this.rows[choice], index: choice}
	}
}

func (this *list) View() string {
	helpText := this.headerStyle.Render(" - Enter: Select   Left: Prev Page   Right: Next Page")
	pageCount := this.headerStyle.Render(fmt.Sprintf("%d/%d", this.page+1, this.pageCount()))
	head := lipgloss.JoinHorizontal(0, pageCount, helpText)

	var body string
	starti := this.page * this.pageSize
	endi := (this.page + 1) * this.pageSize
	if endi >= len(this.rows) {
		endi = len(this.rows)
	}
	for index, row := range this.rows[starti:endi] {
		selected := index == this.pos
		if selected {
			body += this.selectedStyle.Render(" > "+row.name+" ("+row.desc+")") + "\n"
		} else {
			body += this.unselectedStyle.Render("   "+row.name+" ("+row.desc+")") + "\n"
		}
	}

	return lipgloss.JoinVertical(0, head, body[0:len(body)-1])
}
