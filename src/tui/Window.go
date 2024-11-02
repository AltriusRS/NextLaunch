package tui

import (
	"strings"
)

type Window struct {
	Title         string
	width, height int
	border        *Borders
	active        bool
	children      [][]Renderer
}

func (widget *Window) SetWidth(width int) {
	//TODO implement me
	widget.width = width
}

func (widget *Window) SetHeight(height int) {
	//TODO implement me
	widget.height = height
}

func NewWindow(title string, width, height int, active bool) *Window {
	return &Window{
		Title:    title,
		width:    width,
		height:   height,
		border:   NewBorders([4]int{1, 1, 1, 1}, [4]int{1, 1, 1, 1}, title),
		active:   active,
		children: make([][]Renderer, 3),
	}
}

func (widget *Window) SetActive(active bool) {
	widget.active = active
}

func (widget *Window) SetSize(width, height int) {
	widget.width = width
}

func (widget *Window) Size() (int, int) {
	return widget.width, widget.height
}

func (widget *Window) Render(m *Model) string {

	//fmt.Println("Rendering widget")

	lines := make([]string, widget.height)

	if widget.border != nil {
		lines = widget.border.Render(widget.width, widget.height)
	}

	output := strings.Join(lines, "\r\n")

	return output
}

func (widget *Window) Clear() {
	for _, child := range widget.children {
		for _, child := range child {
			child.Clear()
		}
	}
}

func repeat(s string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}
