package tui

import (
	"strings"

	"github.com/gTauro8/git-toxagotchi/assets"
	"github.com/gTauro8/git-toxagotchi/internal/domain"
)

func LoadASCII(stage domain.Stage) string {
	filename := "ascii/" + string(stage) + ".txt"
	data, err := assets.FS.ReadFile(filename)
	if err != nil {
		return "[" + strings.ToUpper(string(stage)) + "]"
	}
	return string(data)
}

func RenderPet(stage domain.Stage, blink bool) string {
	art := LoadASCII(stage)
	if blink {
		art = strings.ReplaceAll(art, "o", "O")
	}
	return art
}
