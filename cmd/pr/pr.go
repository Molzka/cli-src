package pr

import (
	"github.com/spf13/cobra"
)

var PrCmd = &cobra.Command{
	Use:   "pr",
	Short: "Команды для работы с pull requests",
}

func init() {
	PrCmd.AddCommand(listCmd)
	PrCmd.AddCommand(createCmd)
}
