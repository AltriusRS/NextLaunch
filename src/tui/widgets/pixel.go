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
		z_index:  0,
		features: 0,
	}
)

type Pixel struct {
	posX, posY int
	char       string
	z_index    int
	features   byte
	fg         int32
}

func NewPixel(char string, zIndex int, features byte) *Pixel {
	return &Pixel{
		char:     char,
		z_index:  zIndex,
		features: features,
	}
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

	if escape == "\x1b[" {
		return p.char
	} else {
		escape = escape[:len(escape)-1] + "m"
		return fmt.Sprintf("%s%s\x1b[0m", escape, p.char)
	}
}
