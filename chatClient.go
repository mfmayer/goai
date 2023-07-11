package gopenai

import (
	"net/http"
)

type clientOptions struct {
	model     string
	oaiClient *OpenAIClient
}

type clientOption func(co *clientOptions)

func WithModel(model string) clientOption {
	return func(co *clientOptions) {
		co.model = model
	}
}

func WithOpenAIClient(oaiClient *OpenAIClient) clientOption {
	return func(co *clientOptions) {
		co.oaiClient = oaiClient
	}
}

type ChatClient struct {
	*OpenAIClient
	model        string
	openaiApiKey string
	httpClient   *http.Client
}

func NewChatClient(openaiApKey string, opts ...clientOption) *ChatClient {
	options := clientOptions{
		model: "gpt-3.5-turbo",
	}
	if options.oaiClient == nil {
		options.oaiClient = NewOpenAIClient(openaiApKey, "https://api.openai.com")
	}
	for _, opt := range opts {
		opt(&options)
	}
	return &ChatClient{
		OpenAIClient: options.oaiClient,
		openaiApiKey: openaiApKey,
		model:        options.model,
		httpClient:   &http.Client{
			// Timeout: time.Second * 30,
		},
	}
}

func (c *ChatClient) GetChatCompletion(prompt *ChatPrompt) (completion *ChatCompletion, err error) {
	completion = &ChatCompletion{}
	err = c.sendRequest("POST", "/v1/chat/completions", prompt, completion)
	return
}
