package UI

import (
	"math"
	"strings"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type item struct {
	name string
	desc string
}

type list struct {
	page int
	pageSize int
	pos             int
	rows            []item
	selectedStyle   lipgloss.Style
	unselectedStyle lipgloss.Style
}

func newList(items []item, pageS int, s lipgloss.Style, u lipgloss.Style) list {
	return list{pos: 0, rows: items, pageSize: pageS, selectedStyle: s, unselectedStyle: u}
}

func (this *list) pageCount() int {
	return int(math.Ceil(float64(len(this.rows))/float64(this.pageSize)))
}

func (this *list) currentPageSize() int {
	if this.page < this.pageCount()-1 {
		return this.pageSize
	} else {
		return len(this.rows)%this.pageSize
	}
}

func (this *list) PageUp() {
	if this.page > 0 {
		this.page -= 1
	} else {
		this.page = this.pageCount()-1
	}
}

func (this *list) PageDown() {
	if this.page < this.pageCount()-1 {
		this.page += 1
	} else {
		this.page = 0
	}
}

func (this *list) Update(msg tea.Msg) {
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
			this.PageDown()
			if this.pos >= this.currentPageSize() {
				this.pos = this.currentPageSize()-1
			}
		case "right":
			this.PageUp()
			if this.pos >= this.currentPageSize() {
				this.pos = this.currentPageSize()-1
			}
		}
	}
}

func (this *list) View() string {
	head := this.unselectedStyle.Render(fmt.Sprintf("%d/%d", this.page+1, this.pageCount()))
	
	var body string
	starti := this.page*this.pageSize
	endi := (this.page+1)*this.pageSize
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

	return lipgloss.JoinVertical(0, head, body[0:len(body)-1], tail)
}
