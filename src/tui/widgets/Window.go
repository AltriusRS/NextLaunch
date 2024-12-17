package widgets

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

func (widget *Window) Collides(i int, i2 int) bool {
	//TODO implement me
	//panic("implement me")
	return false
}

func (widget *Window) Trigger(kind, x, y int, entity string, data interface{}) {
	//TODO implement me
	panic("implement me")
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
		border:   NewBorders([4]int{1, 1, 1, 1}, title),
		children: []Renderer{},
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
	widget.height = height
}

func (widget *Window) Size() (int, int) {
	return widget.width, widget.height
}

func (widget *Window) Render(width int, height int, focusEntity string) *PixelMap {
	widget.height = height
	widget.width = width

	var pm *PixelMap

	// Render borders if present
	if widget.border != nil {
		pm = widget.border.Render(widget.width, widget.height, widget.z_index)
	}

	if pm == nil {
		pm = NewPixelMap()
	}

	// Render child nodes
	for _, node := range widget.children {
		pm.Ingest(node.Render(widget.width, widget.height, focusEntity))
	}

	return pm
}

func (widget *Window) AddChild(child Renderer) {
	widget.children = append(widget.children, child)
}

func (widget *Window) RemoveChild(id string) {
	for i, child := range widget.children {
		if child.Id() == id {
			widget.children = append(widget.children[:i], widget.children[i+1:]...)
			return
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
