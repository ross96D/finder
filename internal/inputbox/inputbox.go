package inputbox

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/ross96D/finder/internal"
)

type InputBox struct {
	BackgroundColor color.RGBA

	text         internal.Text
	textRelative internal.Position

	rect internal.Rectangle

	IsFocus bool
}

func New(bg color.RGBA, fg color.RGBA, size int32, pos internal.Position) InputBox {
	resp := InputBox{
		BackgroundColor: bg,
		text:            internal.Text{Color: fg, FontSize: size, Value: "Insert some shit"},
		// TODO add width by param
		rect: internal.Rectangle{Pos: pos, Width: 600, Height: size},
	}
	resp.updateTextRelative()
	return resp
}

func (t *InputBox) updateTextRelative() {
	t.text.Position.X = t.rect.Pos.X + t.textRelative.X
	t.text.Position.Y = t.rect.Pos.Y + t.textRelative.Y
}

func (i *InputBox) Rect() rl.Rectangle {
	return i.rect.Rect()
}

func (i *InputBox) Draw() {
	rl.DrawRectangle(i.rect.Pos.X, i.rect.Pos.Y, i.rect.Width, i.rect.Height, i.BackgroundColor)
	i.text.Draw()
}
