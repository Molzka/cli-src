package api

import (
	"bytes"
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

func (c *SourceCraftClient) DoRequest(method, path string, body interface{}) (map[string]interface{}, error) {
	url := c.baseURL + path

	var reqBody io.Reader = nil
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, url, reqBody)
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

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result map[string]interface{}

	if resp.StatusCode != http.StatusOK {
		var apiError APIErrorResponse
		if err := json.Unmarshal(responseBody, &apiError); err == nil {
			if resp.StatusCode == 201 {
				return result, nil
			}
			return nil, fmt.Errorf("API error %d: %s (code: %s)", resp.StatusCode, apiError.Message, apiError.ErrorCode)
		}
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(responseBody))
	}

	if err := json.Unmarshal(responseBody, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return result, nil
}
