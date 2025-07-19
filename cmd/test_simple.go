package main

import (
	"fmt"
	"log"
	"time"

	"github.com/pushkarsingh32/semanticpen-go-sdk"
)

func main() {
	fmt.Println("🚀 Testing Go SDK with Simple Generation")
	fmt.Println("========================================\n")

	// Create client with debug enabled
	client := semanticpen.NewClient("your-api-key-here", &semanticpen.Config{
		Debug: true,
	})

	// Test with only required parameter
	fmt.Println("📝 Generating article with minimal params...")
	response, err := client.GenerateArticle("Simple test article Go SDK", nil)
	if err != nil {
		log.Fatalf("❌ Article generation failed: %v", err)
	}

	fmt.Printf("✅ Generation successful!\n")
	fmt.Printf("   Article ID: %s\n", response.ArticleID)
	fmt.Printf("   Project ID: %s\n", response.ProjectID)
	fmt.Printf("   Message: %s\n", response.Message)

	articleID := response.ArticleID

	// Check status periodically
	fmt.Println("\n⏳ Checking status periodically...")
	maxChecks := 8
	
	for i := 1; i <= maxChecks; i++ {
		time.Sleep(5 * time.Second)

		article, err := client.GetArticle(articleID)
		if err != nil {
			log.Printf("❌ Error checking status: %v", err)
			continue
		}

		fmt.Printf("   Check %d: %s (progress: %d%%)\n", i, article.Status, article.Progress)

		switch article.Status {
		case "finished":
			fmt.Println("   ✅ Article completed successfully!")
			if article.ArticleHTML != "" {
				fmt.Println("   🎉 Has article_html: Yes")
				fmt.Printf("   📏 Content length: %d\n", len(article.ArticleHTML))
			} else {
				fmt.Println("   🎉 Has article_html: No")
			}
			return
		case "failed":
			fmt.Println("   ❌ Article failed")
			if article.ErrorMessage != "" {
				fmt.Printf("   Error details: %s\n", article.ErrorMessage)
			}
			return
		case "pending", "processing":
			if i < maxChecks {
				continue
			}
		}
	}

	fmt.Printf("⏰ Reached maximum checks (%d), but article may still be processing\n", maxChecks)
}