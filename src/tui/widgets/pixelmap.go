package widgets

// PixelMap is a struct to provide useful methods for managing a map of pixels
type PixelMap struct {
	// Provides a map of x/y coordinates to a pixel
	pixels map[int]map[int]Pixel
}

func NewPixelMap() PixelMap {
	return PixelMap{
		pixels: make(map[int]map[int]Pixel),
	}
}

func (m *PixelMap) Ingest(other *PixelMap) {
	for x, row := range other.pixels {
		for y, pixel := range row {
			existing := m.Get(x, y)

			if existing != nil {
				// Only replace the pixel (or add it) if the z-index is higher than the current pixel at the same coordinates
				// This allows us to layer multiple pixels on top of each other
				if pixel.z_index > existing.z_index {
					m.Set(x, y, pixel)
				}
			} else {
				m.Set(x, y, pixel)
			}
		}
	}
}

func (m *PixelMap) Get(x, y int) *Pixel {
	yMap, ok := m.pixels[x]
	if !ok {
		return nil
	}

	pixel, ok := yMap[y]
	if !ok {
		return nil
	}

	return &pixel
}

func (m *PixelMap) Set(x, y int, pixel Pixel) {
	_, ok := m.pixels[x]
	if !ok {
		m.pixels[x] = make(map[int]Pixel)
	}
	m.pixels[x][y] = pixel
}

func (m *PixelMap) IngestRaw(pixels []Pixel) {
	for _, pixel := range pixels {
		existing := m.Get(pixel.posX, pixel.posY)
		if existing != nil {
			if pixel.z_index > existing.z_index {
				m.Set(pixel.posX, pixel.posY, pixel)
			}
		} else {
			m.Set(pixel.posX, pixel.posY, pixel)
		}
	}
}
