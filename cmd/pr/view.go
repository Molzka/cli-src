package pr

import (
	"fmt"
	"src/internal/api"
	"src/internal/utils"

	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view <id>",
	Short: "Посмотреть детали pull request",
	Run:   runView,
}

func runView(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("Ошибка. Введите ID pull request.")
		return
	}

	prId := args[0]

	body, err := api.GetPullRequest(prId)
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

	if body.Author != nil {
		fmt.Printf("│ Author:       %s\n", body.Author["slug"])
	}

	fmt.Printf("│ Created:      %v\n", utils.FormatDate(body.CreatedAt))
	fmt.Printf("│ Updated:      %v\n", utils.FormatDate(body.UpdatedAt))
	if body.Repository != nil {
		fmt.Printf("│ Repository:   %v\n", body.Repository["slug"])
	}
	fmt.Printf("└─────────────────────────────────────────────────────\n\n")
}
