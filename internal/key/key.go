package key

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/ross96D/finder/internal"
)

type Key int32

const (
	None Key = iota + 1000
	CtrlBackspace
	Backspace
	CtrlDelete
	Delete

	MoveWordFoward
	MoveWordBackward
	MoveFoward
	MoveBackward
	MoveUp
	MoveDown
)

// !!!! this model only allows for one consumer
func Poll() Key {
	char := rl.GetCharPressed()
	if char != 0 {
		return Key(char)
	}
	if rl.IsKeyDown(rl.KeyLeftControl) || rl.IsKeyDown(rl.KeyLeftControl) {
		if rl.IsKeyDown(rl.KeyBackspace) {
			return CtrlBackspace
		}
		if rl.IsKeyDown(rl.KeyDelete) {
			return CtrlDelete
		}
		if rl.IsKeyDown(rl.KeyLeft) {
			return MoveWordBackward
		}
		if rl.IsKeyDown(rl.KeyRight) {
			return MoveWordFoward
		}
		return None
	}
	if rl.IsKeyDown(rl.KeyBackspace) {
		return check(Backspace)
	}
	if rl.IsKeyDown(rl.KeyDelete) {
		return check(Delete)
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		return check(MoveBackward)
	}
	if rl.IsKeyDown(rl.KeyRight) {
		return check(MoveFoward)
	}
	if rl.IsKeyDown(rl.KeyUp) {
		return check(MoveBackward)
	}
	if rl.IsKeyDown(rl.KeyDown) {
		return check(MoveFoward)
	}
	return None
}

type prevkey struct {
	Key
	Time internal.AppTime
}

var prev prevkey

// TODO implement scaling velocity
func check(key Key) Key {
	if prev.Key != key {
		prev.Key = key
		prev.Time = internal.Time
		return key
	}
	passedTime := internal.Time - prev.Time
	if passedTime < 10 {
		return None
	}

	prev.Time = internal.Time
	return key
}
