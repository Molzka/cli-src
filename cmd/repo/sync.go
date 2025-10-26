// sync.go
package repo

import (
	"fmt"
	"src/internal/api"
	"src/internal/config"

	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Синхронизировать форк с оригинальным репозиторием",
	Run:   runSync,
}

func runSync(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("Ошибка. Введите название репозитория-форка")
		return
	}

	reponame := args[0]
	err := SyncRepository(reponame)
	if err != nil {
		fmt.Printf("Ошибка синхронизации: %v\n", err)
		return
	}

	fmt.Println("Форк успешно синхронизирован")
}

func SyncRepository(forkRepoName string) error {
	forkRepo, err := api.GetRepoInfo(forkRepoName)
	if err != nil {
		return fmt.Errorf("failed to get fork repository info: %w", err)
	}

	if forkRepo.Parent == nil {
		return fmt.Errorf("репозиторий %s не является форком", forkRepoName)
	}

	fmt.Printf("Найден форк: %s\n", forkRepo.Name)
	fmt.Printf("Родительский репозиторий ID: %s\n", forkRepo.Parent.ID)
	fmt.Printf("Родительский репозиторий Slug: %s\n", forkRepo.Parent.Slug)

	parentRepo, err := api.GetRepoInfo(forkRepo.Organization.Slug + "/" + forkRepo.Slug)
	if err != nil {
		return fmt.Errorf("failed to get parent repository info by ID %s: %w", forkRepo.Parent.ID, err)
	}

	fmt.Printf("Родительский репозиторий: %s/%s\n", parentRepo.Organization.Slug, parentRepo.Slug)
	fmt.Printf("Синхронизация: %s -> %s\n", parentRepo.DefaultBranch, forkRepo.DefaultBranch)

	return CreateSyncPullRequest(forkRepo, parentRepo)
}

func CreateSyncPullRequest(forkRepo, parentRepo *api.Repository) error {
	token, err := config.LoadToken()
	if err != nil {
		return fmt.Errorf("failed to load token: %w", err)
	}

	client := api.NewSourceCraftClient(token)

	path := fmt.Sprintf("/repos/id:%s/pulls", forkRepo.ID)

	prBody := api.CreatePullRequestBody{
		Title:        fmt.Sprintf("Sync with upstream %s", parentRepo.Slug),
		Description:  fmt.Sprintf("Автоматическая синхронизация с оригинальным репозиторием %s\n\nИзменения из ветки `%s` родительского репозитория.", parentRepo.Name, parentRepo.DefaultBranch),
		SourceBranch: parentRepo.DefaultBranch,
		TargetBranch: forkRepo.DefaultBranch,
		ForkRepoID:   parentRepo.ID,
		Publish:      true,
	}

	fmt.Printf("Создание Pull Request...\n")
	fmt.Printf("Источник: %s (репозиторий %s)\n", prBody.SourceBranch, parentRepo.ID)
	fmt.Printf("Цель: %s (репозиторий %s)\n", prBody.TargetBranch, forkRepo.ID)

	_, err = client.DoRequest("POST", path, prBody)
	if err != nil {
		return fmt.Errorf("failed to create sync PR: %w", err)
	}

	return nil
}
