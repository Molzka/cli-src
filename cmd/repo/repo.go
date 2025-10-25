package repo

import (
	"github.com/spf13/cobra"
)

var RepoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Команды для repositories",
}

func init() {
	RepoCmd.AddCommand(listCmd)
}
