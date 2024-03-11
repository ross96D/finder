package inputbox

import (
	"fmt"
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/ross96D/finder/internal"
	"github.com/ross96D/finder/internal/blink"
	"github.com/ross96D/finder/internal/key"
)

type InputBox struct {
	BackgroundColor color.RGBA

	text         internal.Text
	textRelative internal.Position

	rect internal.Rectangle

	isFocus bool
	blink   blink.Blink
}

func New(bg color.RGBA, fg color.RGBA, size int32, pos internal.Position) InputBox {
	blink := blink.New()

	resp := InputBox{
		BackgroundColor: bg,
		text:            internal.Text{Color: fg, FontSize: size, Value: "Insert some shit"},
		// TODO add width by param
		rect:  internal.Rectangle{Pos: pos, Width: 600, Height: size},
		blink: blink,
	}
	resp.updateTextRelative()
	return resp
}

func (t *InputBox) updateTextRelative() {
	t.text.Position.X = t.rect.Pos.X + t.textRelative.X
	t.text.Position.Y = t.rect.Pos.Y + t.textRelative.Y
	t.blink.SetPosition(t.text.Position)

	t.blink.Move(t.text, true)
}

func (i *InputBox) Rect() rl.Rectangle {
	return i.rect.Rect()
}

func (i *InputBox) ChangeFocus() {
	i.isFocus = !i.isFocus
	if i.isFocus {
		i.blink.Init()
	} else {
		i.blink.Blur()
	}
}

func (i *InputBox) Draw() {
	rl.DrawRectangle(i.rect.Pos.X, i.rect.Pos.Y, i.rect.Width, i.rect.Height, i.BackgroundColor)
	i.text.Draw()
	i.blink.Draw()
}

func (i *InputBox) Update() {
	if !i.isFocus {
		return
	}

	keyInput := key.Poll()
keysFor:
	for keyInput != key.None {

		switch keyInput {
		case key.Backspace:
			if i.text.Value == "" {
				break keysFor
			}
			lastChar := i.text.Value[len(i.text.Value)-1]
			i.text.Value = i.text.Value[:len(i.text.Value)-1]
			copy := i.text
			copy.Value = string(rune(lastChar))
			i.blink.Move(copy, false)
			break keysFor

		case key.Delete:
			fmt.Println("Delete is not implemented")
			break keysFor

		case key.CtrlBackspace:
			fmt.Println("CtrlBackspace is not implemented")
			break keysFor

		case key.CtrlDelete:
			fmt.Println("CtrlDelete is not implemented")
			break keysFor

		case key.MoveWordBackward:
			fmt.Println("MoveWordBackward is not implemented")
			break keysFor

		case key.MoveWordFoward:
			fmt.Println("MoveWordFoward is not implemented")
			break keysFor

		case key.MoveBackward:
			fmt.Println("MoveBackward is not implemented")
			break keysFor

		case key.MoveFoward:
			fmt.Println("MoveFoward is not implemented")
			break keysFor

		case key.MoveUp:
			fmt.Println("MoveUp is not implemented")
			break keysFor

		case key.MoveDown:
			fmt.Println("MoveDown is not implemented")
			break keysFor

		case key.None:
			break keysFor

		default:
			if keyInput >= 1 && keyInput <= 125 {
				s := string(rune(keyInput))
				i.text.Value += s
				copy := i.text
				copy.Value = s
				i.blink.Move(copy, true)
			}
		}
		keyInput = key.Poll()
	}
}
