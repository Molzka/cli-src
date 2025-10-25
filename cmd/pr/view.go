package pr

import (
	"bufio"
	"fmt"
	"os"
	"src/internal/api"
	"src/internal/utils"

	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "Получить список pull requests",
	Run:   runView,
}

func runView(cmd *cobra.Command, args []string) {
	fmt.Print("Введите репозиторий (username/repo): ")
	reader := bufio.NewReader(os.Stdin)
	reponame := utils.ReadLine(reader)
	fmt.Print("Введите slug вашего pull requests: ")
	prGluf := utils.ReadLine(reader)
	body, err := api.GetPullRequest(reponame, prGluf)
	if err != nil {
		fmt.Println("Ошибка при получении pull requests: ", err)
		return
	}

	printPullRequest(body)
}

func printPullRequest(body *api.PullRequest) {
	fmt.Printf("┌─────────────────────────────────────────────────────\n")
	fmt.Printf("│ ID:           %v\n", body.ID)
	fmt.Printf("│ Slug:         %v\n", body.Slug)
	fmt.Printf("│ Title:        %v\n", body.Title)
	fmt.Printf("│ Description:  %v\n", body.Description)
	fmt.Printf("│ Status:       %v\n", body.Status)
	fmt.Printf("│ Branch:       %s -> %s\n", body.SourceBranch, body.TargetBranch)
	fmt.Printf("│ Author:       %s\n", body.Author)
	fmt.Printf("│ Created:      %v\n", formatDate(body.CreatedAt))
	fmt.Printf("│ Updated:      %v\n", formatDate(body.UpdatedAt))
	fmt.Printf("│ Repository:   %v\n", body.Repository)
	fmt.Printf("└─────────────────────────────────────────────────────\n\n")
}
