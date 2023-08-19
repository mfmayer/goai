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
		ChatPromptConfig: &gopenai.ChatPromptConfig{
			Model: "gpt-3.5-turbo",
		},
		Messages: []*gopenai.Message{
			{Role: "user", Content: "Hallo"},
		},
	})
	if err != nil {
		t.Fatalf("%v", err)
	}
	if cC.Error != nil {
		t.Fatalf("%v", cC.Error.Message)
	}
	fmt.Println(cC)
}

func TestChatFuinctions(t *testing.T) {
	godotenv.Load()
	key := os.Getenv("OPENAI_API_KEY")
	c := gopenai.NewChatClient(key)
	cC, err := c.GetChatCompletion(&gopenai.ChatPrompt{
		ChatPromptConfig: &gopenai.ChatPromptConfig{
			Model: "gpt-3.5-turbo",
		},
		Messages: []*gopenai.Message{
			{Role: "user", Content: "Wie ist das Wetter in Berlin?"},
		},
		Functions: []*gopenai.Function{
			{
				Name:        "get_weather",
				Description: "Get the weather for a given location",
				Parameters: gopenai.FunctionParameters{
					Type: "object",
					Properties: map[string]gopenai.FunctionParameter{
						"location": {
							Type:        "string",
							Description: "The location to get the weather for",
						},
					},
				},
			},
		},
	})
	if err != nil {
		t.Fatalf("%v", err)
	}
	if cC.Error != nil {
		t.Fatalf("%v", cC.Error.Message)
	}
	fmt.Println(cC)
}
