package tui

import (
	"strings"

	"github.com/gTauro8/git-toxagotchi/assets"
	"github.com/gTauro8/git-toxagotchi/internal/domain"
)

// moodEyes maps mood to the eye characters used in ASCII art.
// The base art uses "o" for eyes; we swap it per mood.
var moodEyes = map[domain.Mood]string{
	domain.MoodEcstatic:  "^",
	domain.MoodHappy:     "^",
	domain.MoodContent:   "o",
	domain.MoodNeutral:   "o",
	domain.MoodAnnoyed:   ">",
	domain.MoodGrumpy:    ">",
	domain.MoodMiserable: "x",
	domain.MoodChaotic:   "@",
}

// moodLabel returns a short mood indicator shown below the pet.
var moodLabel = map[domain.Mood]string{
	domain.MoodEcstatic:  "( ✦ ecstatic ✦ )",
	domain.MoodHappy:     "(  ^ happy ^  )",
	domain.MoodContent:   "(   content   )",
	domain.MoodNeutral:   "(   neutral   )",
	domain.MoodAnnoyed:   "(  > annoyed  )",
	domain.MoodGrumpy:    "( >> grumpy >>)",
	domain.MoodMiserable: "( x_x misery )",
	domain.MoodChaotic:   "( @_@ CHAOS! )",
}

func LoadASCII(stage domain.Stage) string {
	filename := "ascii/" + string(stage) + ".txt"
	data, err := assets.FS.ReadFile(filename)
	if err != nil {
		return "[" + strings.ToUpper(string(stage)) + "]"
	}
	return string(data)
}

func RenderPet(stage domain.Stage, mood domain.Mood, blink bool) string {
	art := LoadASCII(stage)

	eye := moodEyes[mood]
	if eye == "" {
		eye = "o"
	}

	if blink {
		art = strings.ReplaceAll(art, "o", "-")
		art = strings.ReplaceAll(art, eye, "-")
	} else {
		art = strings.ReplaceAll(art, "o", eye)
	}

	label, ok := moodLabel[mood]
	if !ok {
		label = ""
	}
	return art + "\n" + label
}
