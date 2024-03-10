package font

import (
	"github.com/flopp/go-findfont"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var font rl.Font

func Font() rl.Font {
	return font
}

func Load(size int32) {
	s, err := findfont.Find("ComicShannsMonoNerdFontMono-Regular")
	// TODO Craft a better way to load a font
	if err != nil {
		panic(err)
	}
	font = rl.LoadFontEx(s, size, nil, 250)
}

func Unload() {
	rl.UnloadFont(font)
}
