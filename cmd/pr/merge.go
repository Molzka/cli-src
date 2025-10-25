package pr

import (
	"bufio"
	"fmt"
	"os"
	"src/internal/api"
	"src/internal/utils"

	"github.com/spf13/cobra"
)

var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Получить список pull requests",
	Run:   runMerge,
}

func runMerge(cmd *cobra.Command, args []string) {
	fmt.Print("Введите репозиторий (username/repo): ")
	reader := bufio.NewReader(os.Stdin)
	reponame := utils.ReadLine(reader)
	fmt.Print("Введите pr slug: ")
	prGlug := utils.ReadLine(reader)

	resq, err := api.MergePullRequest(reponame, prGlug)

	if err != nil {
		fmt.Println("Ошибка при получении pull requests: ", err)
		return
	}

	fmt.Println(resq)
}
