package widgets

import (
	"sort"
	"strings"
)

// PixelMap is a struct to provide useful methods for managing a map of pixels
type PixelMap struct {
	// Provides a map of y/x coordinates to a pixel
	pixels map[int]map[int]*Pixel

	startX int
	startY int
}

func NewPixelMap() *PixelMap {
	return &PixelMap{
		pixels: make(map[int]map[int]*Pixel),
		startX: 0,
		startY: 0,
	}
}

func (m *PixelMap) SetPos(x, y int) {
	m.startX = x
	m.startY = y
}

func (m *PixelMap) Ingest(other *PixelMap) {

	for y, row := range other.pixels {
		for x, pixel := range row {
			existing := m.Get(x+other.startX, y+other.startY)

			if existing != nil {
				// Only replace the pixel (or add it) if the z-index is higher than the current pixel at the same coordinates
				// This allows us to layer multiple pixels on top of each other
				if pixel.zIndex > existing.zIndex {
					m.Set(x+other.startX, y+other.startY, pixel)
				}
			} else {
				m.Set(x+other.startX, y+other.startY, pixel)
			}
		}
	}
}

func (m *PixelMap) Get(x, y int) *Pixel {
	xMap, ok := m.pixels[y]
	if !ok {
		return nil
	}

	pixel, ok := xMap[x]
	if !ok {
		return nil
	}

	return pixel
}

func (m *PixelMap) Set(x, y int, pixel *Pixel) {
	if m.pixels == nil {
		m.pixels = make(map[int]map[int]*Pixel)
	}

	_, ok := m.pixels[y]
	if !ok {
		m.pixels[y] = make(map[int]*Pixel)
	}
	m.pixels[y][x] = pixel
}

func (m *PixelMap) IngestRaw(pixels []*Pixel) {
	for _, pixel := range pixels {
		existing := m.Get(pixel.posX, pixel.posY)
		if existing != nil {
			if pixel.zIndex > existing.zIndex {
				m.Set(pixel.posX, pixel.posY, pixel)
			}
		} else {
			m.Set(pixel.posX, pixel.posY, pixel)
		}
	}
}

func (m *PixelMap) Render() string {
	var lines []string

	rows := make([]int, 0, len(m.pixels))
	for y := range m.pixels {
		rows = append(rows, y)
	}
	sort.Ints(rows)

	for _, y := range rows {
		line := ""
		row := m.pixels[y]
		columns := make([]int, 0, len(row))
		for x := range row {
			columns = append(columns, x)
		}
		sort.Ints(columns)
		for _, x := range columns {
			pixel := row[x]
			if pixel == nil {
				pixel = &Pixel{char: " ", zIndex: 0, features: 0}
			}
			line += pixel.Render()
		}
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}
