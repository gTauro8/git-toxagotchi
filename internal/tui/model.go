package tui

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/sugar_petauro/git-toxagotchi/internal/application"
	"github.com/sugar_petauro/git-toxagotchi/internal/domain"
)

type tickMsg time.Time

type Model struct {
	pet     *domain.Pet
	service *application.Service
	humor   *application.HumorEngine
	blink   bool
	message string
	events  []string
	width   int
	height  int
}

var (
	titleStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205"))
	barStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("42"))
	labelStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Width(10)
	msgStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("220")).Italic(true)
	helpStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("238"))
)

func NewModel(pet *domain.Pet, svc *application.Service) Model {
	return Model{
		pet:     pet,
		service: svc,
		humor:   application.NewHumorEngine(),
		message: "Welcome! Your pet is ready.",
		events:  []string{},
	}
}

func (m Model) Init() tea.Cmd {
	return tickCmd()
}

func tickCmd() tea.Cmd {
	return tea.Tick(500*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "f":
			if m.pet != nil {
				response, err := m.service.FeedPet(m.pet)
				if err == nil {
					m.message = response
				}
			}
		case "p":
			if m.pet != nil {
				response, err := m.service.PlayWithPet(m.pet)
				if err == nil {
					m.message = response
				}
			}
		case "s":
			if m.pet != nil {
				m.message = fmt.Sprintf("Stage: %s | XP: %d | Mood: %s", m.pet.Stage, m.pet.Experience, m.pet.Mood)
			}
		case "h":
			m.message = "Keys: f=feed, p=play, s=status, q=quit"
		}
	case tickMsg:
		m.blink = !m.blink
		return m, tickCmd()
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

func (m Model) View() string {
	if m.pet == nil {
		return "No pet found. Run: git-toxagotchi init\n"
	}
	return RenderView(m)
}
