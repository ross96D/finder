package window

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/ross96D/finder/internal"
	"github.com/ross96D/finder/internal/font"
	"github.com/ross96D/finder/internal/inputbox"
)

const (
	__width__  = 800
	__heigth__ = 450

	// TODO implement a more dynamic way to set the font size
	FONT_SIZE = 25
)

type Window struct {
	input inputbox.InputBox

	drawFps bool
}

func New() Window {
	input := inputbox.New(rl.LightGray, rl.Black, FONT_SIZE, internal.Position{X: 8, Y: 8})
	input.IsFocus = true
	return Window{input: input, drawFps: true}
}

func (w *Window) Init() {

	rl.SetTargetFPS(60)
	rl.SetConfigFlags(rl.FlagWindowHidden)
	rl.InitWindow(__width__, __heigth__, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	font.Load(FONT_SIZE)
	centerWindow(__width__, __heigth__)

	rl.ClearWindowState(rl.FlagWindowHidden)

	fmt.Printf("Start up time: %s\n", time.Since(internal.StartUpTime).String())

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
		{
			rl.ClearBackground(rl.RayWhite)
			w.Draw()
		}
		rl.EndDrawing()
	}

	font.Unload()
	rl.CloseWindow()
}

func (w *Window) Draw() {
	if w.drawFps {
		rl.DrawFPS(__width__-80, 0)
	}
	w.input.Draw()
}

func centerWindow(windowWidth, windowHeigth int) {
	monitor := rl.GetCurrentMonitor()
	monitorWidth := rl.GetMonitorWidth(monitor)
	monitorHeigth := rl.GetMonitorHeight(monitor)
	rl.SetWindowPosition(monitorWidth/2-windowWidth/2, monitorHeigth/2-windowHeigth/2)
}
