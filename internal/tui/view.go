package tui

import (
	"fmt"
	"strings"
)

func renderBar(label string, value int, max int, color string) string {
	filled := value * 20 / max
	bar := strings.Repeat("█", filled) + strings.Repeat("░", 20-filled)
	styled := barStyle.Render(bar)
	lbl := labelStyle.Render(label)
	return fmt.Sprintf("%s %s %3d%%", lbl, styled, value)
}

func RenderView(m Model) string {
	var b strings.Builder

	// Title
	b.WriteString(titleStyle.Render(fmt.Sprintf("  %s the %s  ", m.pet.Name, m.pet.Species)))
	b.WriteString("\n\n")

	// Pet ASCII art
	art := RenderPet(m.pet.Stage, m.blink)
	b.WriteString(art)
	b.WriteString("\n")

	// Stage and mood
	b.WriteString(fmt.Sprintf("  Stage: %-20s Mood: %s\n", m.pet.Stage, m.pet.Mood))
	b.WriteString(fmt.Sprintf("  XP: %d\n\n", m.pet.Experience))

	// Stats bars
	b.WriteString(renderBar("Energy ", m.pet.Energy, 100, "42"))
	b.WriteString("\n")
	b.WriteString(renderBar("Hunger ", m.pet.Hunger, 100, "196"))
	b.WriteString("\n")
	b.WriteString(renderBar("Stress ", m.pet.Stress, 100, "208"))
	b.WriteString("\n")
	b.WriteString(renderBar("Trust  ", m.pet.Trust, 100, "63"))
	b.WriteString("\n")
	b.WriteString(renderBar("Chaos  ", m.pet.Chaos, 100, "160"))
	b.WriteString("\n\n")

	// Message
	b.WriteString(msgStyle.Render("  " + m.message))
	b.WriteString("\n\n")

	// Help
	b.WriteString(helpStyle.Render("  f=feed  p=play  s=status  h=help  q=quit"))
	b.WriteString("\n")

	return b.String()
}
