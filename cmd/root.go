package cmd

import (
	"src/cmd/auth"
	"src/cmd/issue"
	"src/cmd/pr"
	"src/cmd/repo"

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
	rootCmd.AddCommand(repo.RepoCmd)
	rootCmd.AddCommand(pr.PrCmd)
	rootCmd.AddCommand(issue.IssueCmd)
}
