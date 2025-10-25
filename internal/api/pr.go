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
