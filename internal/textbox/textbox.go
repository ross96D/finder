package textbox

import (
	"fmt"
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/ross96D/finder/internal"
)

// TODO add optional borders
type Textbox struct {
	BackgroundColor color.RGBA

	text         internal.Text
	textRelative internal.Position

	position internal.Position
	Width    int32
	Height   int32
}

func New(pos internal.Position, width int32, heigth int32, text string, textPos internal.Position) Textbox {
	fmt.Println("GRJKSFAS", text)
	resp := Textbox{
		text:         internal.Text{Value: text, FontSize: 20, Color: rl.Black},
		position:     pos,
		textRelative: textPos,
		Width:        width,
		Height:       heigth,
	}
	resp.updateTextRelative()
	return resp
}

func (t *Textbox) SetText(text string) {
	t.text.Value = text
}

func (t *Textbox) Rect() rl.Rectangle {
	return rl.NewRectangle(float32(t.position.X), float32(t.position.Y), float32(t.Width), float32(t.Height))
}

func (t *Textbox) SetTextPosition(pos internal.Position) {
	t.textRelative = pos
}

func (t *Textbox) SetPosition(pos internal.Position) {
	t.position = pos
	t.updateTextRelative()
}

func (t *Textbox) updateTextRelative() {
	t.text.Position.X = t.position.X + t.textRelative.X
	t.text.Position.Y = t.position.Y + t.textRelative.Y
}

func (t Textbox) Draw() {
	rl.DrawRectangle(t.position.X, t.position.Y, t.Width, t.Height, t.BackgroundColor)
	t.text.Draw()
}
