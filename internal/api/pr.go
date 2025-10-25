package api

import (
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
