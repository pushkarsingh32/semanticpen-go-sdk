package semanticpen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	DefaultBaseURL = "https://semanticpen.vercel.app/api"
	DefaultTimeout = 30 * time.Second
)

// Client represents the SemanticPen API client
type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
	debug      bool
}

// Config holds configuration options for the client
type Config struct {
	BaseURL string
	Timeout time.Duration
	Debug   bool
}

// NewClient creates a new SemanticPen client with the given API key and optional config
func NewClient(apiKey string, config *Config) *Client {
	if config == nil {
		config = &Config{}
	}

	if config.BaseURL == "" {
		config.BaseURL = DefaultBaseURL
	}

	if config.Timeout == 0 {
		config.Timeout = DefaultTimeout
	}

	return &Client{
		apiKey:  apiKey,
		baseURL: config.BaseURL,
		debug:   config.Debug,
		httpClient: &http.Client{
			Timeout: config.Timeout,
		},
	}
}

// makeRequest makes an HTTP request to the API
func (c *Client) makeRequest(method, endpoint string, body interface{}) (*http.Response, error) {
	url := c.baseURL + endpoint

	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewBuffer(jsonBody)

		if c.debug {
			fmt.Printf("[DEBUG] Request body: %s\n", string(jsonBody))
		}
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	if c.debug {
		fmt.Printf("[DEBUG] %s %s\n", method, url)
		for k, v := range req.Header {
			if k != "Authorization" {
				fmt.Printf("[DEBUG] %s: %s\n", k, v[0])
			} else {
				fmt.Printf("[DEBUG] %s: Bearer ***\n", k)
			}
		}
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	return resp, nil
}

// TestConnection tests the connection to the API
func (c *Client) TestConnection() error {
	resp, err := c.makeRequest("GET", "/test-connection", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return &APIError{
			StatusCode: resp.StatusCode,
			Message:    string(body),
		}
	}

	return nil
}