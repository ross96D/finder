package blink

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/ross96D/finder/internal"
	"github.com/ross96D/finder/internal/font"
)

type Blink struct {
	blur  bool
	Color color.RGBA
	rect  internal.Rectangle

	time      internal.AppTime
	blinkTime uint64
}

func New() Blink {
	var b Blink
	b.blinkTime = 30
	b.rect.Width = 2
	b.rect.Height = 25
	b.Color = rl.Black
	return b
}

func (b *Blink) Init() {
	b.blur = false
	b.time = internal.Time
}

func (b *Blink) Blur() {
	b.blur = true
}

func (b *Blink) Draw() {
	if b.blinkTime == 0 || b.blur {
		return
	}

	if b.shouldDraw() {
		rl.DrawRectangle(b.rect.Pos.X, b.rect.Pos.Y, b.rect.Width, b.rect.Height, b.Color)
	}
}

func (b *Blink) SetPosition(pos internal.Position) {
	b.rect.Pos = pos
}

func (b *Blink) Move(text internal.Text, positive bool) {
	x, _ := font.MessureText(text.Value, int(text.FontSize))

	if positive {
		b.rect.Pos.X += int32(x)
	} else {
		b.rect.Pos.X -= int32(x)
	}
}

func (b *Blink) shouldDraw() bool {
	timePassed := internal.Time - b.time
	if timePassed < internal.AppTime(b.blinkTime) {
		return true
	}

	if timePassed < internal.AppTime(2*b.blinkTime) {
		return false
	}

	// this could cause issues because we assume that shouldDraw will be called on every tick unless is not active
	// but this seems to be a reasonable assumption
	b.Init()
	return true
}
