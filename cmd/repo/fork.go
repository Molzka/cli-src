package repo

import (
	"fmt"
	"src/internal/api"

	"github.com/spf13/cobra"
)

var forkCmd = &cobra.Command{
	Use:   "fork",
	Short: "Сделать fork репозитория",

	Run: runFork,
}

func runFork(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("Ошибка. Введите название репозитория")
		return
	}

	reponame := args[0]

	api.ForkRepository(reponame)
}
