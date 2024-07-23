/*
Copyright © 2024 anvidev <andreasgylche@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gofast",
	Short: "typeracer terminal game to improve your wpm",
	Long:  `Will get to this`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gofast.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	var lang string
	playCmd.PersistentFlags().StringVarP(&lang, "language", "l", "english", "specify language for generated words")
	// playCmd.Flags().StringVarP(&lang, "language", "l", "english", "specify language for generated words")
}
