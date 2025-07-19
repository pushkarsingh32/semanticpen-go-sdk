# SemanticPen Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/pushkarsingh32/semanticpen-go-sdk.svg)](https://pkg.go.dev/github.com/pushkarsingh32/semanticpen-go-sdk)
[![Go Report Card](https://goreportcard.com/badge/github.com/pushkarsingh32/semanticpen-go-sdk)](https://goreportcard.com/report/github.com/pushkarsingh32/semanticpen-go-sdk)

The official Go SDK for [SemanticPen](https://www.semanticpen.com) - AI Article Writer & SEO Blog Generator. Create high-quality, SEO-optimized articles using advanced AI technology.

## Features

- 🚀 **Simple Integration** - Easy-to-use Go interface
- 🔄 **Automatic Polling** - Built-in article generation status monitoring
- 🛡️ **Error Handling** - Comprehensive error types and validation
- 📊 **Progress Tracking** - Real-time generation progress callbacks
- 🎯 **Type Safety** - Full Go type definitions for all API structures
- 🔧 **Configurable** - Flexible client configuration options

## Installation

```bash
go get github.com/pushkarsingh32/semanticpen-go-sdk
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"
    "time"
    
    "github.com/pushkarsingh32/semanticpen-go-sdk"
)

func main() {
    // Create a new client
    client := semanticpen.NewClient("your-api-key-here", &semanticpen.Config{
        Debug: true,
    })

    // Test connection
    if err := client.TestConnection(); err != nil {
        log.Fatal("Connection failed:", err)
    }

    // Generate an article
    article, err := client.GenerateArticleAndWait(
        "Go Programming Best Practices",
        nil, // Use default options
        &semanticpen.GenerateAndWaitOptions{
            MaxAttempts: 60,
            Interval:    5 * time.Second,
            OnProgress: func(attempt int, status string) {
                fmt.Printf("Attempt %d: %s\\n", attempt, status)
            },
        },
    )
    
    if err != nil {
        log.Fatal("Article generation failed:", err)
    }

    fmt.Printf("✅ Article generated successfully!\\n")
    fmt.Printf("Title: %s\\n", article.Title)
    fmt.Printf("Content length: %d characters\\n", len(article.ArticleHTML))
}
```

## API Reference

### Client Configuration

```go
config := &semanticpen.Config{
    BaseURL: "https://semanticpen.vercel.app/api", // Default
    Timeout: 30 * time.Second,                     // Default
    Debug:   true,                                 // Enable debug logging
}

client := semanticpen.NewClient("your-api-key", config)
```

### Generate Article

```go
// Simple generation
response, err := client.GenerateArticle("Your Target Keyword", nil)

// Advanced generation with options
request := &semanticpen.GenerateArticleRequest{
    TargetKeyword: "AI and Machine Learning",
    Generation: &semanticpen.GenerationOptions{
        ProjectName:    "Tech Blog",
        Language:       "en",
        Country:        "US",
        Perspective:    "first-person",
        Purpose:        "informative",
        ClickbaitLevel: 3,
    },
    SEO: &semanticpen.SEOOptions{
        Title:       "Custom SEO Title",
        Description: "Custom meta description",
        Keywords:    []string{"ai", "machine learning", "technology"},
        UseSchema:   true,
    },
    Writing: &semanticpen.WritingOptions{
        Style:         "professional",
        Tone:          "informative",
        Length:        "long",
        IncludeImages: true,
        ImageStyle:    "modern",
    },
}

response, err := client.GenerateArticle("AI and Machine Learning", request)
```

### Monitor Article Status

```go
// Manual status checking
article, err := client.GetArticle(articleID)
fmt.Printf("Status: %s, Progress: %d%%\\n", article.Status, article.Progress)

// Wait for completion with progress tracking
article, err := client.WaitForArticle(articleID, &semanticpen.GenerateAndWaitOptions{
    MaxAttempts: 60,
    Interval:    5 * time.Second,
    OnProgress: func(attempt int, status string) {
        fmt.Printf("[%s] Attempt %d: %s\\n", 
            time.Now().Format("15:04:05"), attempt, status)
    },
})
```

### Error Handling

```go
article, err := client.GenerateArticle("test", nil)
if err != nil {
    switch e := err.(type) {
    case *semanticpen.APIError:
        fmt.Printf("API Error %d: %s\\n", e.StatusCode, e.Message)
    case *semanticpen.ValidationError:
        fmt.Printf("Validation Error for %s: %s\\n", e.Field, e.Message)
    case *semanticpen.RateLimitError:
        fmt.Printf("Rate Limited: %s\\n", e.Message)
        if e.RetryAfter > 0 {
            fmt.Printf("Retry after %d seconds\\n", e.RetryAfter)
        }
    default:
        fmt.Printf("Unknown error: %s\\n", err)
    }
}
```

## Data Structures

### Article Structure

```go
type Article struct {
    ID           string                 `json:"id"`
    ProjectID    string                 `json:"projectId"`
    Status       string                 `json:"status"`        // "pending", "processing", "finished", "failed"
    Progress     int                    `json:"progress"`      // 0-100
    Title        string                 `json:"title"`
    ArticleHTML  string                 `json:"article_html"`
    ArticleJSON  map[string]interface{} `json:"article_json"`
    SEOData      *SEOData              `json:"seo_data"`
    ErrorMessage string                 `json:"error_message"`
    CreatedAt    time.Time             `json:"created_at"`
    UpdatedAt    time.Time             `json:"updated_at"`
}
```

## Error Types

- **APIError**: HTTP API errors with status codes
- **ValidationError**: Input validation failures  
- **RateLimitError**: Rate limiting with retry information

## Examples

See the `/examples` directory for complete usage examples:

- [Basic Usage](examples/basic/main.go)
- [Advanced Options](examples/advanced/main.go)
- [Progress Tracking](examples/progress/main.go)
- [Error Handling](examples/errors/main.go)

## Requirements

- Go 1.19 or higher
- Valid SemanticPen API key

## Getting an API Key

1. Sign up at [SemanticPen](https://www.semanticpen.com)
2. Navigate to your dashboard
3. Generate an API key in the API section
4. Use the key in your application

## Support

- 📧 Email: contact@semanticpen.com
- 🐛 Issues: [GitHub Issues](https://github.com/pushkarsingh32/semanticpen-go-sdk/issues)
- 📖 Documentation: [SemanticPen API Docs](https://www.semanticpen.com/api-documentation)
- 🏠 Homepage: [SemanticPen](https://www.semanticpen.com)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

---

Built with ❤️ by the SemanticPen team