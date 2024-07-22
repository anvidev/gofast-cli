package game

import (
	"bufio"
	"os"

	"github.com/anvidev/typer/internal/model"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	defaultWidth = 60
)

func StartFromStdin() error {
	var stdin []byte
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		stdin = append(stdin, scanner.Bytes()...)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	text := string(stdin)
	text, err := formatWhitespace(text)
	if err != nil {
		return err
	}

	return start(text)
}

func StartFromRandom() error {
	text := generateWordString(20)

	text, err := formatWhitespace(text)
	if err != nil {
		return err
	}

	return start(text)
}

func start(text string) error {
	game := &model.PlayModel{
		Text: []rune(wrapString(text, defaultWidth)),
	}

	program := tea.NewProgram(game)

	if _, err := program.Run(); err != nil {
		return err
	}
	return nil
}
