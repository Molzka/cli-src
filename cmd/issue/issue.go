package issue

import (
	"github.com/spf13/cobra"
)

func init() {
	IssueCmd.AddCommand(listCmd)
}

var IssueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Команды для списков задач",
}
