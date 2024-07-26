/*
Copyright © 2024 anvidev <andreasgylche@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gofast",
	Short: "typeracer terminal game to improve your wpm",
	Long:  "",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
