package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type SourceCraftClient struct {
	baseURL    string
	token      string
	httpClient *http.Client
}

func NewSourceCraftClient(token string) *SourceCraftClient {
	return &SourceCraftClient{
		baseURL: "https://api.sourcecraft.tech",
		token:   token,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *SourceCraftClient) doRequest(method, path string) ([]byte, error) {
	url := c.baseURL + path

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var apiError APIErrorResponse
		if err := json.Unmarshal(body, &apiError); err == nil {
			return nil, fmt.Errorf("API error %d: %s (code: %s)", resp.StatusCode, apiError.Message, apiError.ErrorCode)
		}
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(body))
	}

	return body, nil
}
