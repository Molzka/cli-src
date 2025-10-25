package api

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"src/internal/config"
	"strings"
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

func CloneRepository(reponame, targetDir string) error {
	repo, err := GetRepoInfo(reponame)
	if err != nil {
		return fmt.Errorf("failed to get repository info: %w", err)
	}

	cloneURL := repo.CloneURL.HTTPS
	if cloneURL == "" {
		return fmt.Errorf("no clone URL available for repository")
	}

	fmt.Printf("Cloning repository %s...\n", reponame)
	fmt.Printf("Clone URL: %s\n", cloneURL)
	fmt.Printf("Target directory: %s\n", targetDir)

	cmd := exec.Command("git", "clone", cloneURL, targetDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to clone repository: %w", err)
	}

	fmt.Printf("Successfully cloned repository to %s\n", targetDir)
	return nil
}

func ForkRepository(reponame string) error {
	_, err := GetRepoInfo(reponame)
	slugs := strings.Split(reponame, "/")
	if err != nil {
		return fmt.Errorf("failed to get repository info: %w", err)
	}

	path := fmt.Sprintf("/repos/%s/fork", reponame)

	token, _ := config.LoadToken()
	client := NewSourceCraftClient(token)

	forkBody := ForkRepositoryBody{
		OrgSlug:           slugs[0],
		Slug:              slugs[1] + "-fork",
		DefaultBranchOnly: false,
	}

	body, err := client.DoRequest("POST", path, forkBody)

	fmt.Println(body)

	return nil
}
