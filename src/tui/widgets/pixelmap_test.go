package widgets

import (
	"testing"
)

func TestPixelMapIngest(t *testing.T) {
	pm := NewPixelMap()
	pm2 := NewPixelMap()
	pixel := NewPixel("=", 0, 0)

	pm2.pixels[0] = map[int]*Pixel{}
	pm2.pixels[0][0] = pixel

	pm.Ingest(pm2)
	if len(pm.pixels) != 1 {
		t.Errorf("Expected 1 pixel, got %d", len(pm.pixels))
	}
	if len(pm.pixels[0]) != 1 {
		t.Errorf("Expected 1 row, got %d", len(pm.pixels[0]))
	}
	if pm.pixels[0][0].zIndex != 0 {
		t.Errorf("Expected pixel with zIndex of 0, got %d", pm.pixels[0][0].zIndex)
	}
}

func TestPixelMapIngestWithLTZIndex(t *testing.T) {
	pm := NewPixelMap()
	pm2 := NewPixelMap()
	pixel := NewPixel("=", 0, 0)
	pixel2 := NewPixel("@", 1, 0)

	pm2.pixels[0] = map[int]*Pixel{}
	pm2.pixels[0][0] = pixel

	pm.pixels[0] = map[int]*Pixel{}
	pm.pixels[0][0] = pixel2

	pm.Ingest(pm2)
	if len(pm.pixels) != 1 {
		t.Errorf("Expected 1 pixel, got %d", len(pm.pixels))
	}
	if len(pm.pixels[0]) != 1 {
		t.Errorf("Expected 1 row, got %d", len(pm.pixels[0]))
	}
	if pm.pixels[0][0].zIndex != 1 {
		t.Errorf("Expected pixel with zIndex of 1, got %d", pm.pixels[0][0].zIndex)
	}
}

func TestPixelMapIngestWithGTZIndex(t *testing.T) {
	pm := NewPixelMap()
	pm2 := NewPixelMap()
	pixel := NewPixel("=", 0, 0)
	pixel2 := NewPixel("@", 1, 0)

	pm2.pixels[0] = map[int]*Pixel{}
	pm2.pixels[0][0] = pixel2

	pm.pixels[0] = map[int]*Pixel{}
	pm.pixels[0][0] = pixel

	pm.Ingest(pm2)
	if len(pm.pixels) != 1 {
		t.Errorf("Expected 1 pixel, got %d", len(pm.pixels))
	}
	if len(pm.pixels[0]) != 1 {
		t.Errorf("Expected 1 row, got %d", len(pm.pixels[0]))
	}
	if pm.pixels[0][0].zIndex != 1 {
		t.Errorf("Expected pixel with zIndex of 1, got %d", pm.pixels[0][0].zIndex)
	}
}

func TestPixelMapIngestRaw(t *testing.T) {
	pm := NewPixelMap()
	pixel := NewPixel("=", 0, 0)
	pm.IngestRaw([]*Pixel{pixel})
	if len(pm.pixels) != 1 {
		t.Errorf("Expected 1 pixel, got %d", len(pm.pixels))
	}
	if len(pm.pixels[0]) != 1 {
		t.Errorf("Expected 1 row, got %d", len(pm.pixels[0]))
	}

	if pm.pixels[0][0].zIndex != 0 {
		t.Errorf("Expected pixel with zIndex of 0, got %d", pm.pixels[0][0].zIndex)
	}
}

func TestPixelMapGet(t *testing.T) {
	pm := NewPixelMap()
	pixel := NewPixel("=", 0, 0)

	pm.pixels[0] = map[int]*Pixel{}
	pm.pixels[0][0] = pixel

	if pm.Get(0, 0) == nil {
		t.Error("Expected a pixel, got nil")
	}
}

func TestPixelMapSet(t *testing.T) {
	pm := NewPixelMap()
	pixel := NewPixel("=", 0, 0)
	pm.Set(0, 0, pixel)

	if pm.Get(0, 0) == nil {
		t.Error("Expected a pixel, got nil")
	}
}

func TestPixelMapRender(t *testing.T) {
	pm := NewPixelMap()
	pixel := NewPixel("=", 0, 0)
	pm.Set(0, 0, pixel)

	rendered := pm.Render()

	if rendered != "=" {
		t.Errorf("Expected rendered string to be '=', got '%s'", rendered)
	}
}
