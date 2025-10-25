package repo

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RepoCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Войти в систему",
	Run:   runRepo,
}

func runRepo(cmd *cobra.Command, args []string) {
	fmt.Println("list")
}
