package tui

type TextBox struct {
	lines    int
	width    int
	height   int
	contents string
	active   bool
	border   *Borders
}

func NewTextBox(lines, width, height int) *TextBox {
	return &TextBox{
		lines:    lines,
		width:    width,
		height:   height,
		contents: "",
		active:   false,
		border:   NewBorders([4]int{1, 1, 1, 1}, [4]int{1, 1, 1, 1}, "Textbox"),
	}
}
func (widget *TextBox) Render(m *Model) string {
	return ""
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

func (widget *TextBox) Size() (int, int) {
	return widget.width, widget.height
}
