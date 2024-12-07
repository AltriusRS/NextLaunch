package widgets

type Borders struct {
	v              int
	pad            [4]int
	margin         [4]int
	title          string
	borderColor    string
	titleColor     string
	showTitle      bool
	horizontalChar string
	verticalChar   string
	cornerCharTop  [2]string
	cornerCharBott [2]string
	titleChar      [2]string
}

var DefaultHorizontalChar = "─"
var DefaultVerticalChar = "│"
var DefaultCornerCharTop = [2]string{"┌", "┐"}
var DefaultCornerCharBottom = [2]string{"└", "┘"}
var DefaultTitleChar = [2]string{"┤ ", " ├"}
var DefaultBorderColor = "\x1b[37m"
var DefaultTitleColor = "\x1b[37m"

func NewBordersWithColor(pad [4]int, margin [4]int, title string, borderColor string, titleColor string) *Borders {
	return &Borders{
		pad:            pad,
		margin:         margin,
		title:          title,
		borderColor:    borderColor,
		titleColor:     titleColor,
		showTitle:      true,
		horizontalChar: DefaultHorizontalChar,
		verticalChar:   DefaultVerticalChar,
		cornerCharTop:  DefaultCornerCharTop,
		cornerCharBott: DefaultCornerCharBottom,
		titleChar:      DefaultTitleChar,
	}
}

func NewBorders(pad [4]int, margin [4]int, title string) *Borders {
	return &Borders{
		pad:            pad,
		margin:         margin,
		title:          title,
		borderColor:    DefaultBorderColor,
		titleColor:     DefaultTitleColor,
		showTitle:      true,
		horizontalChar: DefaultHorizontalChar,
		verticalChar:   DefaultVerticalChar,
		cornerCharTop:  DefaultCornerCharTop,
		cornerCharBott: DefaultCornerCharBottom,
		titleChar:      DefaultTitleChar,
	}
}

//func (border *Borders) RenderTitle(m *Model) string {
//	var s string
//	if border.showTitle {
//		s += fmt.Sprintf("%s%s%s%s%s", border.titleChar[0], border.titleColor, border.title, "\x1b[0m", border.titleChar[1])
//		for i := 0; i < border.pad-len(border.title); i++ {
//			s += " "
//		}
//		s += " "
//	}
//	return s
//}
//
//func (border *Borders) Size() (int, int) {
//	return border.pad, border.pad
//}

func (border *Borders) Left() bool {
	return border.v&1 == 1
}

func (border *Borders) Right() bool {
	return border.v&2 == 2
}

func (border *Borders) Top() bool {
	return border.v&4 == 4
}

func (border *Borders) Bottom() bool {
	return border.v&8 == 8
}

func (border *Borders) All() bool {
	return border.v&15 == 15
}

func (border *Borders) SetLeft(left bool) {
	border.setBit(1, left)
}
func (border *Borders) SetRight(right bool) {
	border.setBit(2, right)
}
func (border *Borders) SetTop(top bool) {
	border.setBit(4, top)
}
func (border *Borders) SetBottom(bottom bool) {
	border.setBit(8, bottom)
}

func (border *Borders) SetAll(all bool) {
	border.setBit(1, all)
	border.setBit(2, all)
	border.setBit(4, all)
	border.setBit(8, all)
}

func (border *Borders) Set(value int) {
	border.v = value
}

func (border *Borders) setBit(bit int, value bool) {
	if value {
		border.v |= 1 << bit
	} else {
		border.v &= ^(1 << bit)
	}
}

//var PaddingPixel = NewPixel(0, 0, " ", 0, 0)

func (border *Borders) Render(width int, height int, zIndex int) PixelMap {

	pm := NewPixelMap()

	//index := 0

	return pm

	//for i := index; i < height; i++ {
	//	if border.pad[0] > 0 {
	//
	//		pm.Set(NewPixel(0, i, strings.Repeat(" ", border.pad[0]), zIndex, 0))
	//	}
	//
	//	pm.IngestRaw()
	//}
	//
	//if border.pad[0] > 0 {
	//
	//}
	//
	//suffix := strings.Repeat(" ", border.pad[1])
	//prefix := strings.Repeat(" ", border.pad[3])
	//
	//titleSize := len(border.title) + len(border.titleChar[0]) + len(border.titleChar[1])
	//maxWidth := width
	//maxTitleRepeat := maxWidth - titleSize
	//
	//for i := index; i < height; i++ {
	//	switch i {
	//	case index:
	//		if border.showTitle && len(border.title) > 0 && titleSize < maxWidth {
	//			lines[i] = prefix + border.cornerCharTop[0] + border.titleChar[0] + border.titleColor + border.title + "\x1b[0m" + border.titleChar[1] + strings.Repeat(border.horizontalChar, maxTitleRepeat) + border.cornerCharTop[1] + suffix
	//		} else {
	//			lines[i] = prefix + border.cornerCharTop[0] + strings.Repeat(border.horizontalChar, width-(+border.pad[1]+border.pad[3])) + border.cornerCharTop[1] + suffix
	//		}
	//	case height - (1 + border.pad[2]):
	//		lines[i] = prefix + border.cornerCharBott[0] + strings.Repeat(border.horizontalChar, width-(2+border.pad[1]+border.pad[3])) + border.cornerCharBott[1] + suffix
	//	default:
	//		if i > height-(1+border.pad[2]+border.margin[2]) {
	//			lines[i] = ""
	//			continue
	//		}
	//		lines[i] = prefix + border.verticalChar + strings.Repeat(" ", width-(2+border.pad[1]+border.pad[3])) + border.verticalChar + suffix
	//	}
	//}
	//
	//return lines
}
