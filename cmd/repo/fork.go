package repo

import (
	"fmt"
	"src/internal/api"

	"github.com/spf13/cobra"
)

var forkCmd = &cobra.Command{
	Use:   "fork",
	Short: "Сделать форк репозитория",

	Run: runFork,
}

func runFork(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("Ошибка. Введите название репозитория")
		return
	}

	reponame := args[0]

	err := api.ForkRepository(reponame)

	if err != nil {
		fmt.Println("Ошибка при создании форка, ", err)
		return
	}

	fmt.Println("Форк успешно создан!")

}
