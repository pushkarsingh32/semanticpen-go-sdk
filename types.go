package semanticpen

import "time"

// GenerateArticleRequest represents the request for generating an article
type GenerateArticleRequest struct {
	TargetKeyword string                 `json:"targetKeyword"`
	Generation    *GenerationOptions     `json:"generation,omitempty"`
	SEO           *SEOOptions           `json:"seo,omitempty"`
	Writing       *WritingOptions       `json:"writing,omitempty"`
	Advanced      map[string]interface{} `json:"advanced,omitempty"`
}

// GenerationOptions contains options for article generation
type GenerationOptions struct {
	ProjectName  string `json:"projectName,omitempty"`
	Language     string `json:"language,omitempty"`
	Country      string `json:"country,omitempty"`
	Perspective  string `json:"perspective,omitempty"`
	Purpose      string `json:"purpose,omitempty"`
	ClickbaitLevel int  `json:"clickbaitLevel,omitempty"`
}

// SEOOptions contains SEO-related options
type SEOOptions struct {
	Title       string   `json:"title,omitempty"`
	Description string   `json:"description,omitempty"`
	Keywords    []string `json:"keywords,omitempty"`
	UseSchema   bool     `json:"useSchema,omitempty"`
}

// WritingOptions contains writing style options
type WritingOptions struct {
	Style         string `json:"style,omitempty"`
	Tone          string `json:"tone,omitempty"`
	Length        string `json:"length,omitempty"`
	IncludeImages bool   `json:"includeImages,omitempty"`
	ImageStyle    string `json:"imageStyle,omitempty"`
}

// GenerateArticleResponse represents the response from article generation
type GenerateArticleResponse struct {
	ArticleID string `json:"articleId"`
	ProjectID string `json:"projectId"`
	Message   string `json:"message"`
}

// Article represents an article with its status and content
type Article struct {
	ID           string                 `json:"id"`
	ProjectID    string                 `json:"projectId"`
	Status       string                 `json:"status"`
	Progress     int                    `json:"progress"`
	Title        string                 `json:"title,omitempty"`
	ArticleHTML  string                 `json:"article_html,omitempty"`
	ArticleJSON  map[string]interface{} `json:"article_json,omitempty"`
	SEOData      *SEOData              `json:"seo_data,omitempty"`
	ErrorMessage string                 `json:"error_message,omitempty"`
	CreatedAt    time.Time             `json:"created_at"`
	UpdatedAt    time.Time             `json:"updated_at"`
}

// SEOData contains SEO information for the article
type SEOData struct {
	Title       string            `json:"title,omitempty"`
	Description string            `json:"description,omitempty"`
	Keywords    []string          `json:"keywords,omitempty"`
	Schema      map[string]interface{} `json:"schema,omitempty"`
}

// GenerateAndWaitOptions contains options for the generate and wait method
type GenerateAndWaitOptions struct {
	MaxAttempts    int                                    `json:"maxAttempts,omitempty"`
	Interval       time.Duration                          `json:"interval,omitempty"`
	OnProgress     func(attempt int, status string)       `json:"-"`
}

// PollingOptions contains options for polling operations
type PollingOptions struct {
	MaxAttempts int           `json:"maxAttempts,omitempty"`
	Interval    time.Duration `json:"interval,omitempty"`
}