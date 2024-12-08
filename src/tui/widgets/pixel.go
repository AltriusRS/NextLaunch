package widgets

import "fmt"

const (
	NONE          byte = iota
	BOLD               = 1
	ITALIC             = 2
	UNDERLINE          = 4
	BLINK              = 8
	REVERSE            = 16
	STRIKETHROUGH      = 32
	DIM                = 64
)

var (
	RESET = Pixel{
		char:     "\x1b[0m",
		zIndex:   0,
		features: 0,
	}
)

type Pixel struct {
	posX, posY int
	char       string
	zIndex     int
	features   byte
	fg         *uint32
	bg         *uint32
	fgStandard *uint8
	bgStandard *uint8
}

func NewPixel(char string, zIndex int, features byte) *Pixel {
	return &Pixel{
		char:     char,
		zIndex:   zIndex,
		features: features,
	}
}

func NewColoredPixel(char string, zIndex int, features byte, fg *uint8, bg *uint8) *Pixel {
	return &Pixel{
		char:       char,
		zIndex:     zIndex,
		features:   features,
		fgStandard: fg,
		bgStandard: bg,
	}
}

func New24BitColoredPixel(char string, zIndex int, features byte, fg *uint32, bg *uint32) *Pixel {
	return &Pixel{
		char:     char,
		zIndex:   zIndex,
		features: features,
		fg:       fg,
		bg:       bg,
	}
}

func RGBToTrueColor(r, g, b uint8) *uint32 {
	decimal := uint32(r) << 16
	decimal |= uint32(g) << 8
	decimal |= uint32(b)

	return &decimal
}

func TrueColorToRGB(color *uint32) (uint8, uint8, uint8) {
	r := uint8((*color >> 16) & 0xFF)
	g := uint8((*color >> 8) & 0xFF)
	b := uint8(*color & 0xFF)

	return r, g, b
}

func (p *Pixel) Render() string {
	escape := "\x1b["

	if p.features&BOLD != 0 {
		escape += "1;"
	}

	if p.features&DIM != 0 {
		escape += "2;"
	}

	if p.features&ITALIC != 0 {
		escape += "3;"
	}

	if p.features&UNDERLINE != 0 {
		escape += "4;"
	}

	if p.features&BLINK != 0 {
		escape += "5;"
	}

	if p.features&REVERSE != 0 {
		escape += "7;"
	}

	if p.features&STRIKETHROUGH != 0 {
		escape += "9;"
	}

	if p.bgStandard != nil {
		escape += fmt.Sprintf("4%d;", p.bgStandard)
	} else if p.bg != nil {
		r, g, b := TrueColorToRGB(p.bg)
		escape += fmt.Sprintf("48;2;%d;%d;%d;", r, g, b)
	}

	if p.fgStandard != nil {
		escape += fmt.Sprintf("3%d;", p.fgStandard)
	} else if p.fg != nil {
		r, g, b := TrueColorToRGB(p.fg)
		escape += fmt.Sprintf("38;2;%d;%d;%d;", r, g, b)
	}

	if escape == "\x1b[" {
		return p.char
	} else {
		escape = escape[:len(escape)-1] + "m"
		return fmt.Sprintf("%s%s\x1b[0m", escape, p.char)
	}
}
