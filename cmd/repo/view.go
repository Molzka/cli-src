package repo

import (
	"fmt"
	"src/internal/api"

	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "Войти в систему",
	Run:   runView,
}

func runView(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("Введите репозиторий (username/repo)")
		return
	}

	response, err := api.GetRepoInfo(args[0])

	if err != nil {
		fmt.Println("Ошибка при получении репозиториев: ", err)
		return
	}

	PrintRepositoryInfo(response)
}

func PrintRepositoryInfo(repo *api.Repository) {
	fmt.Printf("┌─────────────────────────────────────────────────────\n")
	fmt.Printf("│ Repository:   %s\n", repo.Name)
	fmt.Printf("│ Slug:         %s\n", repo.Slug)
	fmt.Printf("│ ID:           %s\n", repo.ID)
	fmt.Printf("│ Organization: %s (%s)\n", repo.Organization.Slug, repo.Organization.ID)
	fmt.Printf("│ Description:  %s\n", repo.Description)
	fmt.Printf("│ Visibility:   %s\n", repo.Visibility)
	fmt.Printf("│ Default Branch: %s\n", repo.DefaultBranch)
	fmt.Printf("│ Last Updated: %s\n", repo.LastUpdated)

	if repo.Language != nil {
		fmt.Printf("│ Language:     %s\n", repo.Language.Name)
	}

	fmt.Printf("│ Counters:\n")
	fmt.Printf("│   - Forks:         %s\n", repo.Counters.Forks)
	fmt.Printf("│   - Pull Requests: %s\n", repo.Counters.PullRequests)
	fmt.Printf("│   - Issues:        %s\n", repo.Counters.Issues)
	fmt.Printf("│   - Tags:          %s\n", repo.Counters.Tags)
	fmt.Printf("│   - Branches:      %s\n", repo.Counters.Branches)

	fmt.Printf("│ Clone URLs:\n")
	fmt.Printf("│   - HTTPS: %s\n", repo.CloneURL.HTTPS)
	fmt.Printf("│   - SSH:   %s\n", repo.CloneURL.SSH)

	if repo.Parent != nil {
		fmt.Printf("│ Forked from: %s (%s)\n", repo.Parent.Slug, repo.Parent.ID)
	}

	fmt.Printf("└─────────────────────────────────────────────────────\n\n")
}
