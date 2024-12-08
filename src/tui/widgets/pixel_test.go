package widgets

import (
	"testing"
)

func TestPixelRenderDefault(t *testing.T) {
	pixel := NewPixel("=", 0, 0)
	output := pixel.Render()
	if output != "=" {
		t.Errorf("Expected '=', got '%s'", output)
	}
}

func TestPixelRenderBold(t *testing.T) {
	pixel := NewPixel("=", 0, BOLD)
	output := pixel.Render()
	if output != "\x1b[1m=\x1b[0m" {
		t.Errorf("Expected '\x1b[1m=\x1b[0m', got '%s'", output)
	}
}

func TestPixelRenderItalic(t *testing.T) {
	pixel := NewPixel("=", 0, ITALIC)
	output := pixel.Render()
	if output != "\x1b[3m=\x1b[0m" {
		t.Errorf("Expected '\x1b[3m=\x1b[0m', got '%s'", output)
	}
}

func TestPixelRenderUnderline(t *testing.T) {
	pixel := NewPixel("=", 0, UNDERLINE)
	output := pixel.Render()
	if output != "\x1b[4m=\x1b[0m" {
		t.Errorf("Expected '\x1b[4m=\x1b[0m', got '%s'", output)
	}
}

func TestPixelRenderBlink(t *testing.T) {
	pixel := NewPixel("=", 0, BLINK)
	output := pixel.Render()
	if output != "\x1b[5m=\x1b[0m" {
		t.Errorf("Expected '\x1b[5m=\x1b[0m', got '%s'", output)
	}
}

func TestPixelRenderReverse(t *testing.T) {
	pixel := NewPixel("=", 0, REVERSE)
	output := pixel.Render()
	if output != "\x1b[7m=\x1b[0m" {
		t.Errorf("Expected '\x1b[7m=\x1b[0m', got '%s'", output)
	}
}

func TestPixelRenderStrikethrough(t *testing.T) {
	pixel := NewPixel("=", 0, STRIKETHROUGH)
	output := pixel.Render()
	if output != "\x1b[9m=\x1b[0m" {
		t.Errorf("Expected '\x1b[9m=\x1b[0m', got '%s'", output)
	}
}

func TestPixelRenderDim(t *testing.T) {
	pixel := NewPixel("=", 0, DIM)
	output := pixel.Render()
	if output != "\x1b[2m=\x1b[0m" {
		t.Errorf("Expected '\x1b[2m=\x1b[0m', got '%s'", output)
	}
}
