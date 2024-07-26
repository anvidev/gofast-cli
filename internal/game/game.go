package game

import (
	"bufio"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func StartStdin() error {
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

func StartRandom(lang string) error {
	text := generateWordString(20, lang)

	text, err := formatWhitespace(text)
	if err != nil {
		return err
	}
	fmt.Println(text)

	return start(text)
}

func start(text string) error {
	game := &PlayModel{
		Text: []rune(wrapString(text, StringWidth)),
	}

	program := tea.NewProgram(game)

	if _, err := program.Run(); err != nil {
		return err
	}
	return nil
}
