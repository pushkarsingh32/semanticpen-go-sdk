package main

import (
	"fmt"
	"log"
	"time"

	"github.com/pushkarsingh32/semanticpen-go-sdk"
)

func main() {
	fmt.Println("🚀 SemanticPen Go SDK - Basic Example")
	fmt.Println("=====================================\n")

	// Create client with debug enabled
	client := semanticpen.NewClient("your-api-key-here", &semanticpen.Config{
		Debug: true,
	})

	// Test connection first
	fmt.Println("🔌 Testing connection...")
	if err := client.TestConnection(); err != nil {
		log.Fatalf("❌ Connection failed: %v", err)
	}
	fmt.Println("✅ Connection successful!\n")

	// Generate article with minimal parameters
	fmt.Println("📝 Generating article...")
	response, err := client.GenerateArticle("Go Programming Best Practices", nil)
	if err != nil {
		log.Fatalf("❌ Article generation failed: %v", err)
	}

	fmt.Printf("✅ Article generation started!\n")
	fmt.Printf("   Article ID: %s\n", response.ArticleID)
	fmt.Printf("   Project ID: %s\n", response.ProjectID)
	fmt.Printf("   Message: %s\n\n", response.Message)

	// Wait for completion with progress tracking
	fmt.Println("⏳ Waiting for article completion...")
	article, err := client.WaitForArticle(response.ArticleID, &semanticpen.GenerateAndWaitOptions{
		MaxAttempts: 60,
		Interval:    5 * time.Second,
		OnProgress: func(attempt int, status string) {
			timestamp := time.Now().Format("15:04:05")
			fmt.Printf("[%s] Attempt %d: %s\n", timestamp, attempt, status)
		},
	})

	if err != nil {
		log.Fatalf("❌ Article completion failed: %v", err)
	}

	// Display results
	fmt.Printf("\n🎉 Article completed successfully!\n")
	fmt.Printf("   Title: %s\n", article.Title)
	fmt.Printf("   Status: %s\n", article.Status)
	fmt.Printf("   Progress: %d%%\n", article.Progress)

	if article.ArticleHTML != "" {
		fmt.Printf("   Content length: %d characters\n", len(article.ArticleHTML))
		fmt.Printf("   Has HTML content: ✅\n")
	} else {
		fmt.Printf("   Content: ❌ No HTML content\n")
	}

	if article.SEOData != nil {
		fmt.Printf("   SEO Title: %s\n", article.SEOData.Title)
		fmt.Printf("   SEO Description: %s\n", article.SEOData.Description)
		fmt.Printf("   Keywords: %v\n", article.SEOData.Keywords)
	}

	fmt.Println("\n🎊 Basic example completed!")
}