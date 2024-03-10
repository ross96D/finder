package window

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/ross96D/finder/internal"
	"github.com/ross96D/finder/internal/textbox"
)

const (
	__width__  = 800
	__heigth__ = 450
)

type Window struct {
	textBox textbox.Textbox
}

func New() Window {
	textBox := textbox.New(internal.Position{X: 32, Y: 32}, 60, 60, "Hello mom", internal.Position{X: 12, Y: 12})
	return Window{textBox: textBox}
}

func (w *Window) Init() {
	rl.SetConfigFlags(rl.FlagWindowHidden)
	rl.InitWindow(__width__, __heigth__, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	centerWindow(__width__, __heigth__)

	rl.ClearWindowState(rl.FlagWindowHidden)

	rl.SetTargetFPS(60)

	w.textBox.Width = __width__/2 - 50
	w.textBox.Height = 180
	w.textBox.BackgroundColor = rl.LightGray
	w.textBox.SetPosition(internal.Position{X: 20, Y: 40})

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
		{
			rl.ClearBackground(rl.RayWhite)
			w.Draw()
		}
		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func (w *Window) Draw() {
	w.textBox.Draw()
}

func centerWindow(windowWidth, windowHeigth int) {
	monitor := rl.GetCurrentMonitor()
	monitorWidth := rl.GetMonitorWidth(monitor)
	monitorHeigth := rl.GetMonitorHeight(monitor)
	rl.SetWindowPosition(monitorWidth/2-windowWidth/2, monitorHeigth/2-windowHeigth/2)
}
