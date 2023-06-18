package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/PullRequestInc/go-gpt3"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatalf("Missing API KEY")
	}
	ctx := context.Background()
	client := gpt3.NewClient(apiKey)

	resp, err := client.Completion(ctx, gpt3.CompletionRequest{
		Prompt:    []string{"the first thing about the show friends is that"},
		Echo:      true,
		MaxTokens: gpt3.IntPtr(30),
		Stop:      []string{"."},
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp.Choices[0].Text)

}
