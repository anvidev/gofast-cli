package model

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	width        = 60
	charsPerWord = 5
)

var (
	centerStyle         = lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center)
	untypedStyle        = lipgloss.NewStyle().Faint(true).Foreground(lipgloss.Color("#D9D9D9"))
	cursorStyle         = untypedStyle.Underline(true)
	typedCorrectStyle   = lipgloss.NewStyle().UnsetFaint()
	typedIncorrectStyle = lipgloss.NewStyle().Background(lipgloss.Color("#FD0000")).Bold(true)
	wpmStyle            = lipgloss.NewStyle().Foreground(lipgloss.Color("#916F10"))
)

type PlayModel struct {
	Text     []rune
	Typed    []rune
	Start    time.Time
	Mistakes int
	Score    float64
	width    int
	height   int
}

func (m PlayModel) Init() tea.Cmd {
	return nil
}

func (m PlayModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" || msg.String() == "esc" {
			return m, tea.Quit
		}

		if len(m.Typed) == len(m.Text) {
			return m, tea.Quit
		}

		if m.Start.IsZero() {
			m.Start = time.Now()
		}

		if msg.String() == "backspace" && len(m.Typed) > 0 {
			m.Typed = m.Typed[:len(m.Typed)-1]
			return m, nil
		}

		char := msg.Runes[0]
		next := rune(m.Text[len(m.Typed)])

		if next == '\n' || next == ' ' {
			m.Typed = append(m.Typed, next)

			if char == ' ' {
				return m, nil
			}
		}

		m.Typed = append(m.Typed, msg.Runes...)

		if char == next {
			m.Score++
		} else {
			m.Mistakes++
		}

	case tea.WindowSizeMsg:
		if msg.Width == 0 && msg.Height == 0 {
			return m, nil
		} else {
			m.width = msg.Width
			m.height = msg.Height
			return m, nil
		}
	}
	return m, nil
}

func (m PlayModel) View() string {
	remaining := m.Text[len(m.Typed):]

	var typed string
	for i, c := range m.Typed {
		if c == rune(m.Text[i]) {
			typed += typedCorrectStyle.Render(string(c))
		} else {
			typed += typedIncorrectStyle.Render(string(m.Text[i]))
		}
	}

	var wpm float64
	if len(m.Typed) > 0 {
		wpm = (m.Score / charsPerWord) / (time.Since(m.Start).Minutes())
	}

	s := fmt.Sprintf(
		"%.2f wpm\n\n%s",
		wpm,
		typed,
	)

	if len(remaining) > 0 {
		s += cursorStyle.Render(string(remaining[:1]))
		s += untypedStyle.Render(string(remaining[1:]))
	}

	game := centerStyle.Height(m.height).Width(m.width).Render(s)

	return game
}
