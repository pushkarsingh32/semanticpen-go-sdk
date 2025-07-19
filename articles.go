package semanticpen

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// GenerateArticle generates a new article with the given target keyword and options
func (c *Client) GenerateArticle(targetKeyword string, options *GenerateArticleRequest) (*GenerateArticleResponse, error) {
	if targetKeyword == "" {
		return nil, &ValidationError{
			Field:   "targetKeyword",
			Message: "target keyword is required",
		}
	}

	request := &GenerateArticleRequest{
		TargetKeyword: targetKeyword,
	}

	if options != nil {
		request.Generation = options.Generation
		request.SEO = options.SEO
		request.Writing = options.Writing
		request.Advanced = options.Advanced
	}

	resp, err := c.makeRequest("POST", "/generate-article", request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if c.debug {
		fmt.Printf("[DEBUG] Response: %s\n", string(body))
	}

	if resp.StatusCode != http.StatusOK {
		return nil, c.parseErrorResponse(resp.StatusCode, body)
	}

	var result GenerateArticleResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

// GetArticle retrieves an article by its ID
func (c *Client) GetArticle(articleID string) (*Article, error) {
	if articleID == "" {
		return nil, &ValidationError{
			Field:   "articleID",
			Message: "article ID is required",
		}
	}

	endpoint := fmt.Sprintf("/articles/%s", articleID)
	resp, err := c.makeRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if c.debug {
		fmt.Printf("[DEBUG] Response: %s\n", string(body))
	}

	if resp.StatusCode != http.StatusOK {
		return nil, c.parseErrorResponse(resp.StatusCode, body)
	}

	var article Article
	if err := json.Unmarshal(body, &article); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &article, nil
}

// DeleteArticle deletes an article by its ID
func (c *Client) DeleteArticle(articleID string) error {
	if articleID == "" {
		return &ValidationError{
			Field:   "articleID",
			Message: "article ID is required",
		}
	}

	endpoint := fmt.Sprintf("/articles/%s", articleID)
	resp, err := c.makeRequest("DELETE", endpoint, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		return c.parseErrorResponse(resp.StatusCode, body)
	}

	return nil
}

// GenerateArticleAndWait generates an article and waits for it to complete
func (c *Client) GenerateArticleAndWait(targetKeyword string, options *GenerateArticleRequest, waitOptions *GenerateAndWaitOptions) (*Article, error) {
	if waitOptions == nil {
		waitOptions = &GenerateAndWaitOptions{
			MaxAttempts: 60,
			Interval:    5 * time.Second,
		}
	}

	if waitOptions.MaxAttempts == 0 {
		waitOptions.MaxAttempts = 60
	}
	if waitOptions.Interval == 0 {
		waitOptions.Interval = 5 * time.Second
	}

	result, err := c.GenerateArticle(targetKeyword, options)
	if err != nil {
		return nil, err
	}

	return c.WaitForArticle(result.ArticleID, waitOptions)
}

// WaitForArticle waits for an article to complete generation
func (c *Client) WaitForArticle(articleID string, options *GenerateAndWaitOptions) (*Article, error) {
	if options == nil {
		options = &GenerateAndWaitOptions{
			MaxAttempts: 60,
			Interval:    5 * time.Second,
		}
	}

	for attempt := 1; attempt <= options.MaxAttempts; attempt++ {
		article, err := c.GetArticle(articleID)
		if err != nil {
			return nil, err
		}

		if options.OnProgress != nil {
			options.OnProgress(attempt, article.Status)
		}

		switch article.Status {
		case "finished":
			return article, nil
		case "failed":
			return nil, fmt.Errorf("article generation failed: %s", article.ErrorMessage)
		case "pending", "processing":
			if attempt < options.MaxAttempts {
				time.Sleep(options.Interval)
				continue
			}
		}
	}

	return nil, fmt.Errorf("article generation timeout after %d attempts", options.MaxAttempts)
}

// parseErrorResponse parses API error responses
func (c *Client) parseErrorResponse(statusCode int, body []byte) error {
	var apiErr APIError
	if err := json.Unmarshal(body, &apiErr); err != nil {
		return &APIError{
			StatusCode: statusCode,
			Message:    string(body),
		}
	}

	apiErr.StatusCode = statusCode
	return &apiErr
}