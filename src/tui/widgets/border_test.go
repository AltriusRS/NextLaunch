package widgets

import (
	"testing"
)

func TestBordersRender(t *testing.T) {
	border := NewBorders([4]int{1, 1, 1, 1}, "Long Title")
	width := 50
	height := 20

	pm := border.Render(width, height, 0)

	if len(pm.pixels) != height {
		t.Errorf("Expected %d columns, got %d", height, len(pm.pixels))
	}

	if len(pm.pixels[0]) != width {
		t.Errorf("Expected %d rows, got %d", width, len(pm.pixels[0]))
	}
}
