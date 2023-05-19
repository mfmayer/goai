package gopenai

import (
	"fmt"
	"os"
	"testing"
)

func TestClient(t *testing.T) {
	key := os.Getenv("OPENAI_API_KEY")

	c := NewChatClient(key)
	cC, err := c.GetChatCompletion(&ChatPrompt{
		Model: "gpt-3.5-turbo",
		Messages: []*Message{
			{Role: "user", Content: "Hallo"},
		},
	})
	if err != nil {
		t.Fatalf("%v", err)
	}
	fmt.Println(cC)
}
