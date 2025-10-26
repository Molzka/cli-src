package issue

import (
	"fmt"
	"src/internal/api"
	"time"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Просмотр списка задач",
	Run:   runList,
}

func runList(cmd *cobra.Command, args []string) {
	var issues *api.ListIssuesAssignedToAuthenticatedUserResponse
	var err error

	if len(args) == 0 {
		issues, err = api.GetIssues()
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		issues, err = api.GetIssuesReponame(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Printf("Found %d issues\n", len(issues.Issues))
	for i, issue := range issues.Issues {
		fmt.Printf("Issue %d:\n", i+1)
		fmt.Printf("  ID: %s\n", issue.ID)
		fmt.Printf("  Title: %s\n", issue.Title)
		fmt.Printf("  Status: %s\n", issue.Status.Name)
		fmt.Printf("  Author: %s\n", issue.Author.Slug)
		fmt.Printf("  Created: %s\n", issue.CreatedAt.Format(time.RFC3339))
	}
}
