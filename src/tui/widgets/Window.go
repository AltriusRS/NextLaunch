package widgets

import (
	"strings"
)

type Window struct {
	id            string
	z_index       int
	title         string
	width, height int
	posX, posY    int
	border        *Borders
	active        bool
	children      []Renderer
}

func (widget *Window) SetWidth(width int) {
	//TODO implement me
	widget.width = width
}

func (widget *Window) SetHeight(height int) {
	//TODO implement me
	widget.height = height
}

func NewWindow(title string, width, height int, zIndex int) *Window {
	return &Window{
		id:       NodeID(),
		z_index:  zIndex,
		title:    title,
		width:    width,
		height:   height,
		border:   NewBorders([4]int{1, 1, 1, 1}, [4]int{1, 1, 1, 1}, title),
		children: make([]Renderer, 2),
	}
}

func (widget *Window) Id() string {
	return widget.id
}

func (widget *Window) ZIndex() int {
	return widget.z_index
}

func (widget *Window) SetZIndex(zIndex int) {
	widget.z_index = zIndex
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

func (widget *Window) Render(width int, height int, focusEntity string) string {

	//fmt.Println("Rendering widget")

	lines := make([]string, widget.height)

	if widget.border != nil {
		lines = widget.border.Render(widget.width, widget.height)
	}

	// Render child nodes

	for _, node := range widget.children {
		output := node.Render(widget.width, widget.height, focusEntity)
		for i, line := range output {
			compositeLine := lines[i+widget.posY]
			lines[i+widget.posY] = compositeLine + line
		}
	}

	output := strings.Join(lines, "\r\n")

	return output
}

func repeat(s string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}
