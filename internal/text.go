package internal

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Position struct{ X, Y int32 }

type Text struct {
	Value    string
	Color    color.RGBA
	FontSize int32
	Position
}

func (t *Text) Draw() {
	rl.DrawText(t.Value, t.X, t.Y, t.FontSize, t.Color)
}
