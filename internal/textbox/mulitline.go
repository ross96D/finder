package textbox

import (
	"image/color"

	"github.com/ross96D/finder/internal"
)

type TextboxMultiline struct {
	BackgroundColor color.RGBA

	text         internal.Text
	textRelative internal.Position

	rect internal.Rectangle
}
