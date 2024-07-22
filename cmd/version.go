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
	Short: "Print the version of Gofast",
	Long:  "Prints the current version installed of Gofast CLI game",
	Run:   versionRun,
}

func versionRun(cmd *cobra.Command, args []string) {
	fmt.Println("gofast version 0.0.2")
}
