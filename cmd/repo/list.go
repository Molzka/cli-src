package repo

import (
	"bufio"
	"fmt"
	"os"
	"src/internal/api"
	"src/internal/utils"

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
	fmt.Print("Введите slug пользователя или организации: ")
	reader := bufio.NewReader(os.Stdin)
	slug := utils.ReadLine(reader)
	response, err := api.GetListRepositories(slug)

	if err != nil {
		fmt.Println("Ошибка при получении репозиториев: ", err)
		return
	}

	PrintRepositoriesSummary(response.Repositories)
}

func PrintRepositoriesSummary(repos []api.Repository) {
	fmt.Printf("Найдено %d Репозиториев:\n\n", len(repos))

	for i, repo := range repos {
		fmt.Printf("%d. %s", i+1, repo.Name)
		if repo.Description != "" {
			fmt.Printf(" - %s", repo.Description)
		}
		fmt.Printf(" [%s]", repo.Visibility)
		fmt.Printf(" (%s форков)", repo.Counters.Forks)
		fmt.Println()
	}
}
