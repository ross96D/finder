package internal

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/ross96D/finder/internal/font"
)

type Position struct{ X, Y int32 }

type Text struct {
	Value    string
	Color    color.RGBA
	FontSize int32
	Position
}

func (t *Text) Draw() {
	rl.DrawTextEx(font.Font(), t.Value, rl.Vector2{X: float32(t.X), Y: float32(t.Y)}, float32(t.FontSize), 0, t.Color)
	// rl.DrawText()
}
