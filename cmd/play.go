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
	Short: "short description for play command",
	Long:  "long description for play command",
	Run:   runPlay,
}

func runPlay(cmd *cobra.Command, args []string) {
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
		err = game.StartStdin()
		break
	default:
		err = game.StartRandom(lang)
	}

	if err != nil {
		fmt.Println("error starting game:", err)
		os.Exit(1)
	}
}

func init() {
	var lang string
	playCmd.PersistentFlags().StringVarP(&lang, "language", "l", "english", "specify language for generated words")
}
