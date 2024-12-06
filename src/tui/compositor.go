package tui

import (
	"Nextlaunch/src/tui/widgets"
)

// Compositor is the main rendering manager for the TUI.
// It manages focus, and renders the widgets to the screen
// in the correct order and position
type Compositor struct {

	// The widgets that are currently being rendered by the compositor
	// The key is the widget's unique ID
	// the value is the widget itself
	// Note: A widget may have its own children, but these are the job of the widget to
	// ensure that they are rendered correctly
	widgets     map[string]widgets.Renderer
	width       int
	height      int
	focusEntity string
}

// NewCompositor creates a new compositor with the given widget as the root widget
func NewCompositor(widget widgets.Renderer) *Compositor {
	width, height := widget.Size()
	compositor := &Compositor{
		widgets:     map[string]widgets.Renderer{widget.Id(): widget},
		width:       width,
		height:      height,
		focusEntity: widget.Id(),
	}
	return compositor

}

/* TODO:
* 1. Switch to IDMap based rendering approach (with z-indexing) instead of slices
     This will allow for more complex (and non-tiled) layouts
     additionally it will allow for more complex focus management, and easier rendering simulation
* 2. Implement focus
* 3. Implement focus cycling


*/

func (compositor *Compositor) Render(width, height int) string {
	output := []string{}
	for id, widget := range compositor.widgets {

	}
}
