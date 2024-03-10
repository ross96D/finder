package textbox

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/ross96D/finder/internal"
)

// TODO add optional borders
type Textbox struct {
	BackgroundColor color.RGBA

	text         internal.Text
	textRelative internal.Position

	rect internal.Rectangle
}

func New(pos internal.Position, width int32, heigth int32, text string, textPos internal.Position) Textbox {
	resp := Textbox{
		text:         internal.Text{Value: text, FontSize: 20, Color: rl.Black},
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

func (t *Textbox) SetText(text string) {
	t.text.Value = text
}

func (t *Textbox) Rect() rl.Rectangle {
	return t.rect.Rect()
}

func (t *Textbox) SetTextPosition(pos internal.Position) {
	t.textRelative = pos
}

func (t *Textbox) SetPosition(pos internal.Position) {
	t.rect.Pos = pos
	t.updateTextRelative()
}

func (t *Textbox) updateTextRelative() {
	t.text.Position.X = t.rect.Pos.X + t.textRelative.X
	t.text.Position.Y = t.rect.Pos.Y + t.textRelative.Y
}

func (t Textbox) Draw() {
	rl.DrawRectangle(t.rect.Pos.X, t.rect.Pos.Y, t.rect.Width, t.rect.Height, t.BackgroundColor)
	t.text.Draw()
}
