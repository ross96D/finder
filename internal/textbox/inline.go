package textbox

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/ross96D/finder/internal"
)

// TODO add optional borders
type TextboxInline struct {
	BackgroundColor color.RGBA

	text         internal.Text
	textRelative internal.Position

	rect internal.Rectangle
}

func NewInline(pos internal.Position, width int32, heigth int32, text string, textPos internal.Position) TextboxInline {
	resp := TextboxInline{
		text:         internal.Text{Value: text, FontSize: 25, Color: rl.Black},
		textRelative: textPos,
		rect: internal.Rectangle{
			Pos:    pos,
			Width:  width,
			Height: heigth,
		},
	}
	resp.updateTextRelative()
	return resp
}

func (t *TextboxInline) SetText(text string) {
	t.text.Value = text
}

func (t *TextboxInline) Rect() rl.Rectangle {
	return t.rect.Rect()
}

func (t *TextboxInline) SetTextPosition(pos internal.Position) {
	t.textRelative = pos
}

func (t *TextboxInline) SetPosition(pos internal.Position) {
	t.rect.Pos = pos
	t.updateTextRelative()
}

func (t *TextboxInline) updateTextRelative() {
	t.text.Position.X = t.rect.Pos.X + t.textRelative.X
	t.text.Position.Y = t.rect.Pos.Y + t.textRelative.Y
}

func (t TextboxInline) Draw() {
	rl.DrawRectangle(t.rect.Pos.X, t.rect.Pos.Y, t.rect.Width, t.rect.Height, t.BackgroundColor)
	t.text.Draw()
}
