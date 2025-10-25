package api

import (
	"encoding/json"
	"fmt"
	"src/internal/config"
)

func GetListRepositories(orgSlug string) (*ListRepositoriesResponse, error) {
	path := fmt.Sprintf("/orgs/%s/repos", orgSlug)

	token, _ := config.LoadToken()
	client := NewSourceCraftClient(token)

	body, err := client.DoRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var response ListRepositoriesResponse

	jsonData, _ := json.Marshal(body)
	err = json.Unmarshal(jsonData, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func GetRepoInfo(reponame string) (*Repository, error) {
	path := fmt.Sprintf("/repos/%s", reponame)

	token, _ := config.LoadToken()
	client := NewSourceCraftClient(token)

	body, err := client.DoRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var response Repository

	jsonData, _ := json.Marshal(body)
	err = json.Unmarshal(jsonData, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
