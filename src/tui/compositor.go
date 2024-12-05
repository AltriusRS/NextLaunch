package tui

import (
	"Nextlaunch/src/tui/widgets"
)

type Compositor struct {
	widgets     []widgets.Renderer
	width       int
	height      int
	focusEntity string
}

func NewCompositor(screen widgets.Renderer) *Compositor {
	compositor := &Compositor{
		screen:    screen,
		widgets:   make([]widgets.Renderer, 0),
		active:    0,
		width:     screen.Size()[0],
		height:    screen.Size()[1],
		focused:   0,
		lastFocus: 0,
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
