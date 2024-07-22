package cmd

import (
	"fmt"
	"os"

	"github.com/anvidev/gofast/internal/game"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(playCmd)
}

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "start a Gofast game",
	Long:  "start a Gofast game to practice your touchtyping and improve your wpm",
	Run:   playRun,
}

func playRun(cmd *cobra.Command, args []string) {
	stat, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch true {
	case (stat.Mode() & os.ModeCharDevice) == 0:
		err = game.StartFromStdin()
		break
	default:
		err = game.StartFromRandom()
	}

	if err != nil {
		fmt.Println("error starting game:", err)
		os.Exit(1)
	}
}
