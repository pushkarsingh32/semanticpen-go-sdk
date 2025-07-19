package main

import (
	"fmt"

	"github.com/pushkarsingh32/semanticpen-go-sdk"
)

func main() {
	fmt.Println("🔌 Testing SemanticPen Go SDK Connection")
	fmt.Println("========================================\n")

	// Create client with debug enabled
	client := semanticpen.NewClient("your-api-key-here", &semanticpen.Config{
		Debug: true,
	})

	// Test the connection endpoint directly
	fmt.Println("📡 Testing connection with debug...")
	err := client.TestConnection()
	if err != nil {
		fmt.Printf("❌ Connection test failed: %v\n", err)
		
		// Print error details if available
		switch e := err.(type) {
		case *semanticpen.APIError:
			fmt.Printf("   Status Code: %d\n", e.StatusCode)
			fmt.Printf("   Message: %s\n", e.Message)
			if e.Details != "" {
				fmt.Printf("   Details: %s\n", e.Details)
			}
		}
	} else {
		fmt.Println("✅ Connection test successful!")
	}

	fmt.Println("\n📡 Testing with actual article generation as connection test...")
	
	// If we can generate an article, the connection is good
	response, err := client.GenerateArticle("Connection test Go SDK", nil)
	if err != nil {
		fmt.Printf("❌ Real API test failed: %v\n", err)
		return
	}

	fmt.Println("✅ Connection is working (article generation successful)")
	articleID, _ := response.GetArticleID()
	fmt.Printf("   Article ID: %s\n", articleID)
	fmt.Printf("   Project ID: %s\n", response.ProjectID)
	fmt.Printf("   Message: %s\n", response.Message)
}