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
	Long: `start a Gofast game to practice your touchtyping and improve your wpm
          
          use the language flag to change the language. supported languages are:
            - english
            - danish
            - dutch
            - croatian
            - french
            - georgian
            - german
            - italian
            - norwegian
            - polish
            - portuguese
            - spanish
            - swedish`,
	Run: playRun,
}

func playRun(cmd *cobra.Command, args []string) {
	lang, err := cmd.Flags().GetString("language")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

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
		err = game.StartFromRandom(lang)
	}

	if err != nil {
		fmt.Println("error starting game:", err)
		os.Exit(1)
	}
}
