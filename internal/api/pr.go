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

func GetPullRequest(prId string) (*PullRequest, error) {
	token, _ := config.LoadToken()
	client := NewSourceCraftClient(token)

	path := fmt.Sprintf("/pulls/id:%s", prId)

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

func PublishPullRequest(prId string) error {
	path := fmt.Sprintf("/pulls/id:%s/publish", prId)
	token, _ := config.LoadToken()
	client := NewSourceCraftClient(token)

	resp, err := client.DoRequest("POST", path, nil)
	if err != nil {
		return err
	}

	fmt.Println(resp)

	return nil
}

func MergePullRequest(prId string) (*PullRequest, error) {
	pr, err := GetPullRequest(prId)
	if err != nil {
		return nil, fmt.Errorf("failed to get PR: %w", err)
	}

	fmt.Printf("Current PR status: %s\n", pr.Status)

	if pr.Status == PRStatusDraft {
		fmt.Println("Publishing draft PR...")
		if err := PublishPullRequest(prId); err != nil {
			return nil, fmt.Errorf("failed to publish PR: %w", err)
		}
	}

	fmt.Println("Approving PR...")
	_, err = UpdatePullRequestDecision(prId, ReviewDecisionApprove)
	if err != nil {
		return nil, fmt.Errorf("failed to approve PR: %w", err)
	}

	fmt.Printf("PR successfully moved to merging status\n")
	return nil, nil
}

func UpdatePullRequestDecision(prId string, decision ReviewDecision) (*PullRequest, error) {
	path := fmt.Sprintf("/pulls/id:%s/decision", prId)
	token, _ := config.LoadToken()
	client := NewSourceCraftClient(token)

	decisionReq := struct {
		ReviewDecision ReviewDecision `json:"review_decision"`
	}{
		ReviewDecision: decision,
	}

	body, err := client.DoRequest("POST", path, decisionReq)
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
