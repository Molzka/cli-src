package api

import (
	"encoding/json"
	"fmt"
	"src/internal/config"
)

func GetPrList(reponame string) (map[string]interface{}, error) {
	token, _ := config.LoadToken()
	client := NewSourceCraftClient(token)

	slug := fmt.Sprintf("/repos/%s/pulls", reponame)

	body, err := client.DoRequest("GET", slug, nil)

	return body, err
}

func GetPullRequest(reponame, prSlug string) (*PullRequest, error) {
	token, _ := config.LoadToken()
	client := NewSourceCraftClient(token)

	path := fmt.Sprintf("/repos/%s/pulls/%s", reponame, prSlug)

	body, err := client.DoRequest("GET", path, nil)

	if err != nil {
		return nil, err
	}

	var pr PullRequest
	jsonData, _ := json.Marshal(body)
	err = json.Unmarshal(jsonData, &pr)
	if err != nil {
		return nil, err
	}

	return &pr, nil
}

func PublishPullRequest(reponame, prSlug string) error {
	path := fmt.Sprintf("/repos/%s/pulls/%s/publish", reponame, prSlug)
	token, _ := config.LoadToken()
	client := NewSourceCraftClient(token)

	resp, err := client.DoRequest("POST", path, nil)
	if err != nil {
		return err
	}

	fmt.Println(resp)

	return nil
}

func UpdatePullRequestStatus(reponame, prSlug string, status PRStatus) (*PullRequest, error) {
	path := fmt.Sprintf("/repos/%s/pulls/%s", reponame, prSlug)
	token, _ := config.LoadToken()
	client := NewSourceCraftClient(token)

	updateReq := UpdatePullRequestRequest{
		Status: status,
	}

	body, err := client.DoRequest("PATCH", path, updateReq)
	if err != nil {
		return nil, err
	}

	var pr PullRequest
	jsonData, _ := json.Marshal(body)
	err = json.Unmarshal(jsonData, &pr)
	if err != nil {
		return nil, err
	}

	return &pr, nil
}

func MergePullRequest(reponame, prSlug string) (*PullRequest, error) {
	pr, err := GetPullRequest(reponame, prSlug)
	if err != nil {
		return nil, fmt.Errorf("failed to get PR: %w", err)
	}

	fmt.Printf("Current PR status: %s\n", pr.Status)

	if pr.Status == PRStatusDraft {
		fmt.Println("Publishing draft PR...")
		if err := PublishPullRequest(reponame, prSlug); err != nil {
			return nil, fmt.Errorf("failed to publish PR: %w", err)
		}
		pr, err = GetPullRequest(reponame, prSlug)
		if err != nil {
			return nil, fmt.Errorf("failed to get PR after publishing: %w", err)
		}
	}

	if pr.Status != PRStatusOpen {
		return nil, fmt.Errorf("PR is not in open status, current status: %s", pr.Status)
	}

	fmt.Println("Initiating merge...")
	mergedPR, err := UpdatePullRequestStatus(reponame, prSlug, PRStatusMerging)
	if err != nil {
		return nil, fmt.Errorf("failed to update PR status to merging: %w", err)
	}

	fmt.Printf("PR successfully moved to merging status\n")

	if mergedPR.MergeInfo != nil {
		if mergedPR.MergeInfo.Error != "" {
			return nil, fmt.Errorf("merge failed: %s", mergedPR.MergeInfo.Error)
		}
		if mergedPR.MergeInfo.MergeCommitHash != "" {
			fmt.Printf("Merge completed successfully. Merge commit: %s\n", mergedPR.MergeInfo.MergeCommitHash)
		}
	}

	return mergedPR, nil
}
