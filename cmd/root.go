package cmd

import (
	"src/cmd/auth"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "src",
	Short: "CLI инструмент для работы с Sourcecraft",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(auth.AuthCmd)
}
