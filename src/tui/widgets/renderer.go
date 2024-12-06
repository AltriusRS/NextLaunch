package widgets

import (
	"crypto/rand"
	"encoding/hex"
)

// Renderer is a general interface which all widgets must provide so they can be rendered appropriately
type Renderer interface {
	// Id returns the unique identifier of the widget in the renderer namespace
	Id() string

	// Render takes a width, height, and active id parameter in that order
	// and returns an array of strings representing the rendered final widget
	Render(int, int, string) []string

	// SetWidth sets the width of the widget
	SetWidth(width int)

	// SetHeight sets the height of the widget
	SetHeight(height int)

	// Size returns the width and height of the widget
	Size() (int, int)

	// ZIndex returns the z-index of the widget
	ZIndex() int

	// SetZIndex sets the z-index of the widget
	SetZIndex(zIndex int)

	// Collides caluclates if an interaction at X, Y, collides with the widget
	// Note: This should not account for z-index
	// Params:
	// - x, y: the coordinates of the interaction
	// Returns: bool - true if there is a collision, false otherwise
	Collides(int, int) bool

	// Trigger is the handler for system triggers like mouse clicks or key presses.
	// Check the specific trigger type for more information
	// Params:
	// - kind: The type of trigger
	// - x, y: The coordinates of the trigger
	// - entity: The entity that is "active" at the time of the trigger (if this is a mouse click, this will be the clicked widget at the highest z-index)
	// - data: Any additional data associated with the trigger (eg, lmb click data, or key press data)
	// Returns: nothing - This is a pass-through method that is called by the compositor.
	Trigger(kind, x, y int, entity string, data interface{})
}

func NodeID() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}
