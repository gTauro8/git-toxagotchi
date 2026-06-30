package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	appconfig "github.com/sugar_petauro/git-toxagotchi/internal/infrastructure/config"
)

// ConfigWizard is a step-by-step Bubble Tea setup wizard.
// It collects pet name, theme, and hook preference, then saves config.

type wizardStep int

const (
	stepWelcome wizardStep = iota
	stepPetName
	stepTheme
	stepHook
	stepDone
)

type ConfigWizardModel struct {
	step     wizardStep
	cfg      *appconfig.Config
	input    string
	themeIdx int
	hookIdx  int
	err      error
	Finished bool
	FinalCfg *appconfig.Config
}

var themes = []string{"default", "dracula", "nord", "solarized"}
var hookChoices = []string{"yes — block dangerous commits", "no  — warnings only"}

func NewConfigWizard(cfg *appconfig.Config) ConfigWizardModel {
	themeIdx := 0
	for i, t := range themes {
		if t == cfg.Theme {
			themeIdx = i
			break
		}
	}
	hookIdx := 0
	if !cfg.HookBlocking {
		hookIdx = 1
	}
	return ConfigWizardModel{
		step:     stepWelcome,
		cfg:      cfg,
		input:    cfg.PetName,
		themeIdx: themeIdx,
		hookIdx:  hookIdx,
	}
}

func (m ConfigWizardModel) Init() tea.Cmd {
	return nil
}

func (m ConfigWizardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.step {
		case stepWelcome:
			if msg.Type == tea.KeyEnter || msg.String() == " " {
				m.step = stepPetName
			}
			if msg.Type == tea.KeyCtrlC {
				return m, tea.Quit
			}

		case stepPetName:
			switch msg.Type {
			case tea.KeyEnter:
				if strings.TrimSpace(m.input) != "" {
					m.cfg.PetName = strings.TrimSpace(m.input)
				}
				m.step = stepTheme
			case tea.KeyBackspace, tea.KeyDelete:
				if len(m.input) > 0 {
					m.input = m.input[:len(m.input)-1]
				}
			case tea.KeyCtrlC:
				return m, tea.Quit
			default:
				if msg.Type == tea.KeyRunes {
					m.input += string(msg.Runes)
				}
			}

		case stepTheme:
			switch msg.String() {
			case "left", "h":
				m.themeIdx = (m.themeIdx - 1 + len(themes)) % len(themes)
			case "right", "l":
				m.themeIdx = (m.themeIdx + 1) % len(themes)
			case "enter", " ":
				m.cfg.Theme = themes[m.themeIdx]
				m.step = stepHook
			case "ctrl+c":
				return m, tea.Quit
			}

		case stepHook:
			switch msg.String() {
			case "left", "h":
				m.hookIdx = (m.hookIdx - 1 + len(hookChoices)) % len(hookChoices)
			case "right", "l":
				m.hookIdx = (m.hookIdx + 1) % len(hookChoices)
			case "enter", " ":
				m.cfg.HookBlocking = m.hookIdx == 0
				if err := appconfig.Save(m.cfg); err != nil {
					m.err = err
				}
				m.FinalCfg = m.cfg
				m.Finished = true
				m.step = stepDone
			case "ctrl+c":
				return m, tea.Quit
			}

		case stepDone:
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m ConfigWizardModel) View() string {
	title := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205")).
		Render

	subtitle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("241")).
		Render

	selected := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("86")).
		Padding(0, 1).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("86")).
		Render

	normal := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Padding(0, 1).
		Render

	hint := lipgloss.NewStyle().
		Foreground(lipgloss.Color("238")).
		Render

	var b strings.Builder

	switch m.step {
	case stepWelcome:
		b.WriteString(title("🐣 Welcome to Git-Toxagotchi!"))
		b.WriteString("\n\n")
		b.WriteString("  Your terminal pet companion awaits.\n")
		b.WriteString("  It watches your Git habits and reacts accordingly.\n\n")
		b.WriteString(subtitle("  Let's get you set up in a few steps.\n\n"))
		b.WriteString(hint("  Press ENTER to continue"))

	case stepPetName:
		b.WriteString(title("Step 1 / 3 — Name your pet"))
		b.WriteString("\n\n")
		b.WriteString("  What should we call it?\n\n")
		inputStyle := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("205")).
			Padding(0, 1).
			Width(30)
		b.WriteString("  " + inputStyle.Render(m.input+"▌"))
		b.WriteString("\n\n")
		b.WriteString(hint("  Type a name, then press ENTER"))

	case stepTheme:
		b.WriteString(title("Step 2 / 3 — Choose a theme"))
		b.WriteString("\n\n")
		b.WriteString("  How should Git-Toxagotchi look?\n\n  ")
		for i, t := range themes {
			if i == m.themeIdx {
				b.WriteString(selected(t))
			} else {
				b.WriteString(normal(t))
			}
			b.WriteString(" ")
		}
		b.WriteString("\n\n")
		b.WriteString(hint("  ← → to choose, ENTER to confirm"))

	case stepHook:
		b.WriteString(title("Step 3 / 3 — Pre-commit hook"))
		b.WriteString("\n\n")
		b.WriteString("  Should the hook block commits that look dangerous?\n")
		b.WriteString(subtitle("  (e.g. staged .env files or detected secrets)\n\n  "))
		for i, c := range hookChoices {
			if i == m.hookIdx {
				b.WriteString(selected(c))
			} else {
				b.WriteString(normal(c))
			}
			b.WriteString(" ")
		}
		b.WriteString("\n\n")
		b.WriteString(hint("  ← → to choose, ENTER to confirm"))

	case stepDone:
		b.WriteString(title("✅ All set!"))
		b.WriteString("\n\n")
		b.WriteString(fmt.Sprintf("  Pet name : %s\n", m.cfg.PetName))
		b.WriteString(fmt.Sprintf("  Theme    : %s\n", m.cfg.Theme))
		blocking := "yes"
		if !m.cfg.HookBlocking {
			blocking = "no"
		}
		b.WriteString(fmt.Sprintf("  Blocking : %s\n\n", blocking))
		b.WriteString(subtitle(fmt.Sprintf("  Config saved to %s\n\n", appconfig.Path())))
		b.WriteString(hint("  Press ENTER to continue"))
	}

	return lipgloss.NewStyle().Margin(1, 2).Render(b.String()) + "\n"
}

// RunWizard runs the config wizard and returns the resulting config.
// Returns nil if the user cancelled.
func RunWizard(cfg *appconfig.Config) (*appconfig.Config, error) {
	m := NewConfigWizard(cfg)
	p := tea.NewProgram(m, tea.WithAltScreen())
	result, err := p.Run()
	if err != nil {
		return nil, err
	}
	final := result.(ConfigWizardModel)
	if !final.Finished {
		return nil, nil
	}
	return final.FinalCfg, nil
}
