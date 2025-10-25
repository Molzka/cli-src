package pr

import (
	"bufio"
	"fmt"
	"os"
	"src/internal/api"
	"src/internal/utils"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Получить список pull requests",
	Run:   runList,
}

func runList(cmd *cobra.Command, args []string) {
	fmt.Print("Введите репозиторий (username/repo): ")
	reader := bufio.NewReader(os.Stdin)
	reponame := utils.ReadLine(reader)
	body, err := api.GetPrList(reponame)
	if err != nil {
		fmt.Println("Ошибка при получении pull requests: ", err)
		return
	}

	printPullRequests(body)
}

func printPullRequests(body map[string]interface{}) {
	prs, ok := body["pull_requests"].([]interface{})
	if !ok {
		fmt.Println("No pull requests found")
		return
	}

	fmt.Printf("найдено %d pull requests:\n\n", len(prs))

	for i, pr := range prs {
		prMap := pr.(map[string]interface{})

		fmt.Printf("┌─── PR #%d ──────────────────────────────────────────\n", i+1)
		fmt.Printf("│ ID:           %v\n", prMap["id"])
		fmt.Printf("│ Slug:         %v\n", prMap["slug"])
		fmt.Printf("│ Title:        %v\n", prMap["title"])
		fmt.Printf("│ Status:       %v\n", prMap["status"])
		fmt.Printf("│ Branch:       %s -> %s\n", prMap["source_branch"], prMap["target_branch"])

		if author, ok := prMap["author"].(map[string]interface{}); ok {
			fmt.Printf("│ Author:       %s\n", author["slug"])
		}

		fmt.Printf("│ Created:      %v\n", utils.FormatDate(prMap["created_at"].(string)))
		fmt.Printf("│ Updated:      %v\n", utils.FormatDate(prMap["updated_at"].(string)))
		fmt.Printf("└─────────────────────────────────────────────────────\n\n")
	}
}
