package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "task",
	Short: "A simple CLI task tracker",
}

var filePath string;

func init() {
	rootCommand.PersistentFlags().StringVarP(&filePath, "file", "f", "tasks.json", "use with many json files")
}

func Execute() {
	err := rootCommand.Execute()
	if err != nil {
		os.Exit(1)
	}
}
