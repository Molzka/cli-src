package api

import (
	"encoding/json"
	"fmt"
	"src/internal/config"
)

func GetListRepositories(orgSlug string) (*ListOrganizationRepositoriesResponse, error) {
	path := fmt.Sprintf("/orgs/%s/repos", orgSlug)

	token, _ := config.LoadToken()
	client := NewSourceCraftClient(token)

	body, err := client.DoRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var response ListOrganizationRepositoriesResponse

	jsonData, _ := json.Marshal(body)
	err = json.Unmarshal(jsonData, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
