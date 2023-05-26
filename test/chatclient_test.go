package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/mfmayer/gopenai"
)

func TestChatClient(t *testing.T) {
	godotenv.Load()
	key := os.Getenv("OPENAI_API_KEY")

	c := gopenai.NewChatClient(key)
	cC, err := c.GetChatCompletion(&gopenai.ChatPrompt{
		Model: "gpt-3.5-turbo",
		Messages: []*gopenai.Message{
			{Role: "user", Content: "Hallo"},
		},
	})
	if err != nil {
		t.Fatalf("%v", err)
	}
	fmt.Println(cC)
}
