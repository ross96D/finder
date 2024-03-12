package window

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/ross96D/finder/internal"
	"github.com/ross96D/finder/internal/font"
	"github.com/ross96D/finder/internal/inputbox"
	"github.com/ross96D/finder/internal/key"
	"github.com/ross96D/finder/internal/textbox"
)

const (
	__width__  = 800
	__heigth__ = 450

	// TODO implement a more dynamic way to set the font size
	FONT_SIZE = 25
)

type Window struct {
	input   inputbox.InputBox
	matches []textbox.TextboxInline

	drawFps bool
}

var App Window

func New() *Window {
	App = Window{drawFps: true}
	return &App
}

func (w *Window) Init() {

	rl.SetTargetFPS(60)
	rl.SetConfigFlags(rl.FlagWindowHidden)
	rl.InitWindow(__width__, __heigth__, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	font.Load(FONT_SIZE)
	input := inputbox.New(rl.LightGray, rl.Black, FONT_SIZE, internal.Position{X: 8, Y: 8})
	input.ChangeFocus()
	w.input = input

	centerWindow(__width__, __heigth__)

	rl.ClearWindowState(rl.FlagWindowHidden)

	fmt.Printf("Start up time: %s\n", time.Since(internal.StartUpTime).String())

	for !rl.WindowShouldClose() {
		// time counter for the application..
		// for simplicity sake the time will be the frames
		// (
		//	this can cause issues if the frames drop..
		// 	but maybe we can craft a good solution even with this handicap
		// )
		internal.Time++

		w.Update()

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

func (w *Window) Update() {
	keyInput := key.Poll()
keysFor:
	for keyInput != key.None {
		switch keyInput {
		case key.Backspace:
			w.input.Update(keyInput)

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

		case key.Enter:
			w.appendMatch()
			w.matches[len(w.matches)-1].SetText(w.input.Text())

		case key.None:
			break keysFor

		default:
			w.input.Update(keyInput)
		}
		keyInput = key.Poll()
	}

	if w.matches != nil {
		for _, match := range w.matches {
			match.Draw()
		}
	}
}

func (w *Window) appendMatch() {
	if w.matches == nil {
		w.matches = make([]textbox.TextboxInline, 0)
	}
	var match textbox.TextboxInline
	var rect rl.Rectangle
	if len(w.matches) == 0 {
		rect = w.input.Rect()
	} else {
		rect = w.matches[len(w.matches)-1].Rect()
	}
	match = textbox.NewInline(internal.Position{X: rect.ToInt32().X, Y: rect.ToInt32().Y + rect.ToInt32().Height + 10}, rect.ToInt32().Width/2, rect.ToInt32().Height, w.input.Text(), internal.Position{X: 2, Y: 2})
	w.matches = append(w.matches, match)
}

func centerWindow(windowWidth, windowHeigth int) {
	monitor := rl.GetCurrentMonitor()
	monitorWidth := rl.GetMonitorWidth(monitor)
	monitorHeigth := rl.GetMonitorHeight(monitor)
	rl.SetWindowPosition(monitorWidth/2-windowWidth/2, monitorHeigth/2-windowHeigth/2)
}
