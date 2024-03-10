package internal

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var StartUpTime time.Time

type Rectangle struct {
	Pos    Position
	Width  int32
	Height int32
}

func (r *Rectangle) Rect() rl.Rectangle {
	return rl.NewRectangle(float32(r.Pos.X), float32(r.Pos.Y), float32(r.Width), float32(r.Height))
}
