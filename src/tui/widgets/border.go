package widgets

type Borders struct {
	v              int
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
var DefaultTitleChar = [2]string{"┤", "├"}
var DefaultBorderColor = "\x1b[37m"
var DefaultTitleColor = "\x1b[37m"

func NewBordersWithColor(margin [4]int, title string, borderColor string, titleColor string) *Borders {
	return &Borders{
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

func NewBorders(margin [4]int, title string) *Borders {
	return &Borders{
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

func isWithinMargin(x, y, leftLimit, rightLimit, topLimit, bottomLimit int) bool {
	isTopMargin := y < topLimit
	isBottomMargin := y > bottomLimit
	isLeftMargin := x < leftLimit
	isRightMargin := x > rightLimit
	outcome := isTopMargin || isBottomMargin || isLeftMargin || isRightMargin
	return outcome
}

func (border *Borders) Render(width int, height int, zIndex int) *PixelMap {
	pm := NewPixelMap()

	paddingPixel := NewPixel(" ", zIndex, 0)

	// Margins follow CSS convention (Top, Right, Bottom, Left)
	leftLimit := border.margin[3]
	rightLimit := width - 1 - border.margin[1]
	topLimit := border.margin[0]
	bottomLimit := height - 1 - border.margin[2]

	leftWall := leftLimit
	rightWall := width - 1 - border.margin[1]
	topWall := topLimit
	bottomWall := height - 1 - border.margin[2]

	titleWidth := len(border.title) + 2

	titleStart := leftWall + 2
	titleEnd := titleStart + titleWidth

	// Render margin pixels
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			//	Set the top and bottom margin
			if isWithinMargin(x, y, leftLimit, rightLimit, topLimit, bottomLimit) {
				// Set the top margin
				pm.Set(x, y, paddingPixel)
			} else {
				// Set the pixel as a blank
				pm.Set(x, y, paddingPixel)
			}

			// Render border walls
			if y == topWall && x == leftWall {
				// If the pixel is in the top left corner, render the top left corner character
				pm.Set(x, y, NewPixel(border.cornerCharTop[0], zIndex, 0))
			} else if y == topWall && x == rightWall {
				// If the pixel is in the top right corner, render the top right corner character
				pm.Set(x, y, NewPixel(border.cornerCharTop[1], zIndex, 0))
			} else if y == bottomWall && x == leftWall {
				// If the pixel is in the bottom left corner, render the bottom left corner character
				pm.Set(x, y, NewPixel(border.cornerCharBott[0], zIndex, 0))
			} else if y == bottomWall && x == rightWall {
				// If the pixel is in the bottom right corner, render the bottom right corner character
				pm.Set(x, y, NewPixel(border.cornerCharBott[1], zIndex, 0))
			} else if (x == leftWall || x == rightWall) && y > topWall && y < bottomWall {
				// If the pixel is in either vertical surface, render the vertical wall character
				pm.Set(x, y, NewPixel(border.verticalChar, zIndex, 0))
			} else if (y == topWall || y == bottomWall) && x > leftWall && x < rightWall {
				// If the pixel is in either horizontal surface, render the horizontal wall character
				pm.Set(x, y, NewPixel(border.horizontalChar, zIndex, 0))
			}

			// If enabled, render the title
			if border.showTitle && y == topWall {
				if x == titleStart {
					pm.Set(x, y, NewPixel(border.titleChar[0], zIndex, 0))
				} else if x == titleEnd+1 {
					pm.Set(x, y, NewPixel(border.titleChar[1], zIndex, 0))
				} else if x == titleStart+1 || x == titleEnd {
					pm.Set(x, y, NewPixel(" ", zIndex, 0))
				} else {
					if x >= titleStart+1 && x <= titleEnd {
						index := x - titleStart - 2
						pm.Set(x, y, NewPixel(border.title[index:index+1], zIndex, 0))
					}
				}
			}
		}
	}

	//// Render walls
	//for y := 0; y < height; y++ {
	//	for x := 0; x < width; x++ {
	//
	//	}
	//}
	//
	//// Render title

	return pm
}
