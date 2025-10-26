package pr

import (
	"fmt"
	"src/internal/api"

	"github.com/spf13/cobra"
)

var mergeCmd = &cobra.Command{
	Use:   "merge <id>",
	Short: "Сделать merge pull request'a",
	Run:   runMerge,
}

func runMerge(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("Ошибка. Введите ID pull request.")
		return
	}

	prId := args[0]

	api.MergePullRequest(prId)
}
