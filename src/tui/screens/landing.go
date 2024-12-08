package screens

import (
	"Nextlaunch/src/tui/widgets"
)

func LandingScreen(ctx RenderContext) widgets.Renderer {
	window := widgets.NewWindow("Landing", 80, 20, 1)

	return window
}
