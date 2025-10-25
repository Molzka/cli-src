package pr

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Получить список pull requests",
	Run:   runList,
}

func runList(cmd *cobra.Command, args []string) {
	fmt.Println("list")
}
