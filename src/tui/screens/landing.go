package screens

import (
	"Nextlaunch/src/tui/widgets"
)

func LandingScreen(ctx RenderContext) widgets.Renderer {
	window := widgets.NewWindow("Landing", ctx.Width, ctx.Height, 1)

	leftPane := widgets.NewWindow("Left Pane", ctx.Width/2, ctx.Height, 2)
	leftPane.SetActive(true)

	rightPane := widgets.NewWindow("Right Pane", ctx.Width/2, ctx.Height, 3)
	rightPane.SetActive(false)

	window.AddChild(leftPane)
	window.AddChild(rightPane)

	//leftPane.AddChild(widgets.NewTextBox(1, ctx.Width/2, ctx.Height/2, 4, nil))

	return window
}
