package game

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	StringWidth  = 60
	charsPerWord = 5
)

var (
	center    = lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center)
	untyped   = lipgloss.NewStyle().Faint(true).Foreground(lipgloss.Color("#D9D9D9"))
	cursor    = untyped.Underline(true)
	correct   = lipgloss.NewStyle().UnsetFaint()
	incorrect = lipgloss.NewStyle().Background(lipgloss.Color("#FF6961")).Bold(true)
)

type TickMsg time.Time

func doTick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

type PlayModel struct {
	Text      []rune
	Typed     []rune
	start     time.Time
	Mistakes  int
	Score     float64
	width     int
	height    int
	countdown int
}

func (m PlayModel) Init() tea.Cmd {
	return doTick()
}

func (m PlayModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEscape:
			return m, tea.Quit
		}

		if len(m.Typed) == len(m.Text) {
			return m, tea.Quit
		}

		if m.start.IsZero() {
			m.start = time.Now()
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

	case TickMsg:
		m.countdown++
		return m, doTick()
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
			typed += correct.Render(string(c))
		} else {
			typed += incorrect.Render(string(m.Text[i]))
		}
	}

	var wpm float64
	if len(m.Typed) > 0 {
		wpm = (m.Score / charsPerWord) / (time.Since(m.start).Minutes())
	}

	s := fmt.Sprintf(
		"%.0f wpm  %s seconds left\n\n%s",
		wpm,
		fmt.Sprint(m.countdown),
		typed,
	)

	if len(remaining) > 0 {
		s += cursor.Render(string(remaining[:1]))
		s += untyped.Render(string(remaining[1:]))
	}

	textBox := lipgloss.NewStyle().MaxWidth(StringWidth)
	game := center.Height(m.height).Width(m.width).Render(textBox.Render(s))

	return game
}
