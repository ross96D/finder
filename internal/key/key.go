package key

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/ross96D/finder/internal"
)

type Key int32

func (k Key) isUp() bool {
	switch k {
	case None:
		return false
	case CtrlBackspace:
		return (rl.IsKeyUp(rl.KeyLeftControl) || rl.IsKeyUp(rl.KeyRightControl)) && rl.IsKeyUp(rl.KeyBackspace)
	case Backspace:
		return rl.IsKeyUp(rl.KeyBackspace)
	case CtrlDelete:
		return (rl.IsKeyUp(rl.KeyLeftControl) || rl.IsKeyUp(rl.KeyRightControl)) && rl.IsKeyUp(rl.KeyDelete)
	case Delete:
		return rl.IsKeyUp(rl.KeyDelete)
	case Enter:
		return rl.IsKeyUp(rl.KeyEnter)
	case MoveWordFoward:
		return (rl.IsKeyUp(rl.KeyLeftControl) || rl.IsKeyUp(rl.KeyRightControl)) && rl.IsKeyUp(rl.KeyRight)
	case MoveWordBackward:
		return (rl.IsKeyUp(rl.KeyLeftControl) || rl.IsKeyUp(rl.KeyRightControl)) && rl.IsKeyUp(rl.KeyLeft)
	case MoveFoward:
		return rl.IsKeyUp(rl.KeyRight)
	case MoveBackward:
		return rl.IsKeyUp(rl.KeyLeft)
	case MoveUp:
		return rl.IsKeyUp(rl.KeyUp)
	case MoveDown:
		return rl.IsKeyUp(rl.KeyDown)
	default:
		return rl.IsKeyUp(int32(k))
	}
}

const (
	None Key = iota + 1000
	CtrlBackspace
	Backspace
	CtrlDelete
	Delete

	Enter

	MoveWordFoward
	MoveWordBackward
	MoveFoward
	MoveBackward
	MoveUp
	MoveDown
)

// !!!! this model only allows for one consumer
func Poll() Key {
	if prev.repeats > 0 && prev.isUp() {
		prev.repeats = 0
	}

	char := rl.GetCharPressed()
	if char != 0 {
		return check(Key(char))
	}
	if rl.IsKeyDown(rl.KeyLeftControl) || rl.IsKeyDown(rl.KeyLeftControl) {
		if rl.IsKeyDown(rl.KeyBackspace) {
			return check(CtrlBackspace)
		}
		if rl.IsKeyDown(rl.KeyDelete) {
			return check(CtrlDelete)
		}
		if rl.IsKeyDown(rl.KeyLeft) {
			return check(MoveWordBackward)
		}
		if rl.IsKeyDown(rl.KeyRight) {
			return check(MoveWordFoward)
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
		return check(MoveUp)
	}
	if rl.IsKeyDown(rl.KeyDown) {
		return check(MoveDown)
	}
	if rl.IsKeyDown(rl.KeyEnter) {
		return check(Enter)
	}
	return None
}

type prevkey struct {
	Key
	Time    internal.AppTime
	repeats int
}

func (p prevkey) waitTime() internal.AppTime {
	switch p.repeats {
	case 0:
		return 10
	case 1:
		return 10
	case 2:
		return 5
	default:
		return 0
	}
}

var prev prevkey

// TODO implement scaling velocity
func check(key Key) Key {
	if prev.Key != key {
		prev.Key = key
		prev.Time = internal.Time
		prev.repeats = 0
		return key
	}
	passedTime := internal.Time - prev.Time
	if passedTime <= prev.waitTime() {
		return None
	}
	prev.repeats += 1
	prev.Time = internal.Time
	return key
}
