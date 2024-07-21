package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of Typer",
	Long:  "Prints the current version installed of Typer CLI game",
	Run:   versionRun,
}

func versionRun(cmd *cobra.Command, args []string) {
	fmt.Println("typer version 0.0.1")
}
