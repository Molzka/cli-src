package pr

import (
	"github.com/spf13/cobra"
)

type PullRequest struct {
	ID           string            `json:"id"`
	Slug         string            `json:"slug"`
	Title        string            `json:"title"`
	Description  string            `json:"description"`
	Status       string            `json:"status"`
	SourceBranch string            `json:"source_branch"`
	TargetBranch string            `json:"target_branch"`
	Author       map[string]string `json:"author"`
	CreatedAt    string            `json:"created_at"`
	UpdatedAt    string            `json:"updated_at"`
	UpdatedBy    map[string]string `json:"updated_by"`
	Repository   map[string]string `json:"repository"`
}

var PrCmd = &cobra.Command{
	Use:   "pr",
	Short: "Команды для работы с pull requests",
}

func init() {
	PrCmd.AddCommand(listCmd)
	PrCmd.AddCommand(createCmd)
}
