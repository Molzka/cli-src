package repo

import (
	"fmt"
	"src/internal/api"

	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Склонировать репозиторий на устройство",

	Run: runClone,
}

func runClone(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("Ошибка. Введите название репозитория")
		return
	}

	reponame := args[0]

	api.CloneRepository(reponame, reponame)
}
