package api

import (
	"encoding/json"
	"fmt"
	"src/internal/config"
)

func GetIssues() (*ListIssuesAssignedToAuthenticatedUserResponse, error) {
	token, _ := config.LoadToken()
	client := NewSourceCraftClient(token)
	body, _ := client.DoRequest("GET", "/me/issues", nil)

	var response ListIssuesAssignedToAuthenticatedUserResponse
	jsonData, _ := json.Marshal(body)

	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func GetIssuesReponame(reponame string) (*ListIssuesAssignedToAuthenticatedUserResponse, error) {
	token, _ := config.LoadToken()
	client := NewSourceCraftClient(token)
	path := fmt.Sprintf("/repos/%s/issues", reponame)
	body, _ := client.DoRequest("GET", path, nil)

	var response ListIssuesAssignedToAuthenticatedUserResponse
	jsonData, _ := json.Marshal(body)

	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func CreateIssue(repoInfo *Repository, issueData *CreateIssueBody) {
	token, _ := config.LoadToken()
	client := NewSourceCraftClient(token)
	path := fmt.Sprintf("/repos/id:%s/issues", repoInfo.ID)

	body, _ := client.DoRequest("POST", path, issueData)

	fmt.Println(body)
}
