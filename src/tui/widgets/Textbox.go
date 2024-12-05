package widgets

import (
	"strings"
)

type TextBox struct {
	Id       string
	lines    int
	width    int
	height   int
	contents string
	active   bool
	blurred  bool
	border   *Borders
}

func NewTextBox(lines, width, height int) *TextBox {
	return &TextBox{
		Id:       NodeID(),
		lines:    lines,
		width:    width,
		height:   height,
		contents: "",
		active:   false,
		blurred:  false,
		border:   NewBorders([4]int{1, 1, 1, 1}, [4]int{1, 1, 1, 1}, "Textbox"),
	}
}
func (widget *TextBox) Render() string {
	if widget.blurred {
		return strings.Repeat("*", len(widget.contents))
	}
	return widget.contents
}

func (widget *TextBox) Clear() {

}

func (widget *TextBox) SetWidth(width int) {
	widget.width = width
}

func (widget *TextBox) SetHeight(height int) {
	widget.height = height
}

func (widget *TextBox) SetActive(active bool) {
	widget.active = active
}

func (widget *TextBox) SetBorder(border *Borders) {
	widget.border = border
}

func (widget *TextBox) SetContents(contents string) {
	widget.contents = contents
}

func (widget *TextBox) Size() (int, int) {
	return widget.width, widget.height
}

func (widget *TextBox) SetFocus(state bool) {
	widget.active = state
}

func (widget *TextBox) SetBlur(state bool) {
	widget.blurred = state
}

func (widget *TextBox) IsActive() bool {
	return widget.active
}
