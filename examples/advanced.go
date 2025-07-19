package main

import (
	"fmt"
	"log"
	"time"

	"github.com/pushkarsingh32/semanticpen-go-sdk"
)

func main() {
	fmt.Println("🚀 SemanticPen Go SDK - Advanced Example")
	fmt.Println("=========================================\n")

	// Create client with custom configuration
	client := semanticpen.NewClient("your-api-key-here", &semanticpen.Config{
		Debug:   true,
		Timeout: 60 * time.Second,
	})

	// Test connection
	fmt.Println("🔌 Testing connection...")
	if err := client.TestConnection(); err != nil {
		log.Fatalf("❌ Connection failed: %v", err)
	}
	fmt.Println("✅ Connection successful!\n")

	// Generate article with advanced options
	fmt.Println("📝 Generating article with advanced options...")

	request := &semanticpen.GenerateArticleRequest{
		TargetKeyword: "Artificial Intelligence in Healthcare",
		Generation: &semanticpen.GenerationOptions{
			ProjectName:    "Healthcare Tech Blog",
			Language:       "en",
			Country:        "US",
			Perspective:    "third-person",
			Purpose:        "informative",
			ClickbaitLevel: 2,
		},
		SEO: &semanticpen.SEOOptions{
			Title:       "AI Revolution in Healthcare: Transforming Patient Care",
			Description: "Discover how artificial intelligence is revolutionizing healthcare, improving patient outcomes, and transforming medical practices worldwide.",
			Keywords:    []string{"artificial intelligence", "healthcare", "medical AI", "patient care", "healthcare technology"},
			UseSchema:   true,
		},
		Writing: &semanticpen.WritingOptions{
			Style:         "professional",
			Tone:          "informative",
			Length:        "long",
			IncludeImages: true,
			ImageStyle:    "professional",
		},
		Advanced: map[string]interface{}{
			"includeStatistics": true,
			"includeCaseStudies": true,
			"targetAudience":     "healthcare professionals",
		},
	}

	response, err := client.GenerateArticle("Artificial Intelligence in Healthcare", request)
	if err != nil {
		log.Fatalf("❌ Article generation failed: %v", err)
	}

	fmt.Printf("✅ Article generation started!\n")
	fmt.Printf("   Article ID: %s\n", response.ArticleID)
	fmt.Printf("   Project ID: %s\n", response.ProjectID)
	fmt.Printf("   Message: %s\n\n", response.Message)

	// Wait for completion with detailed progress tracking
	fmt.Println("⏳ Waiting for article completion with detailed tracking...")
	
	startTime := time.Now()
	article, err := client.WaitForArticle(response.ArticleID, &semanticpen.GenerateAndWaitOptions{
		MaxAttempts: 100,
		Interval:    3 * time.Second,
		OnProgress: func(attempt int, status string) {
			elapsed := time.Since(startTime).Round(time.Second)
			timestamp := time.Now().Format("15:04:05")
			fmt.Printf("[%s] Attempt %d (%v elapsed): %s\n", timestamp, attempt, elapsed, status)
		},
	})

	if err != nil {
		log.Fatalf("❌ Article completion failed: %v", err)
	}

	// Display comprehensive results
	fmt.Printf("\n🎉 Article completed successfully!\n")
	fmt.Printf("==========================================\n")
	fmt.Printf("📋 Article Details:\n")
	fmt.Printf("   ID: %s\n", article.ID)
	fmt.Printf("   Project ID: %s\n", article.ProjectID)
	fmt.Printf("   Title: %s\n", article.Title)
	fmt.Printf("   Status: %s\n", article.Status)
	fmt.Printf("   Progress: %d%%\n", article.Progress)
	fmt.Printf("   Created: %s\n", article.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Printf("   Updated: %s\n", article.UpdatedAt.Format("2006-01-02 15:04:05"))

	fmt.Printf("\n📄 Content Analysis:\n")
	if article.ArticleHTML != "" {
		fmt.Printf("   HTML Content: ✅ (%d characters)\n", len(article.ArticleHTML))
	} else {
		fmt.Printf("   HTML Content: ❌\n")
	}

	if article.ArticleJSON != nil && len(article.ArticleJSON) > 0 {
		fmt.Printf("   JSON Content: ✅ (%d fields)\n", len(article.ArticleJSON))
	} else {
		fmt.Printf("   JSON Content: ❌\n")
	}

	fmt.Printf("\n🎯 SEO Analysis:\n")
	if article.SEOData != nil {
		fmt.Printf("   Title: %s\n", article.SEOData.Title)
		fmt.Printf("   Description: %s\n", article.SEOData.Description)
		fmt.Printf("   Keywords: %v\n", article.SEOData.Keywords)
		if article.SEOData.Schema != nil && len(article.SEOData.Schema) > 0 {
			fmt.Printf("   Schema: ✅ (%d properties)\n", len(article.SEOData.Schema))
		} else {
			fmt.Printf("   Schema: ❌\n")
		}
	} else {
		fmt.Printf("   No SEO data available\n")
	}

	totalTime := time.Since(startTime).Round(time.Second)
	fmt.Printf("\n⏱️  Total Generation Time: %v\n", totalTime)
	fmt.Println("\n🎊 Advanced example completed!")
}