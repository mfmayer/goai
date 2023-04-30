package oai

// Message represents a message object in a chat conversation.
// Role is the sender's role, either "system", "user", or "assistant".
// Content is the text content of the message.
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatPrompt represents a chat prompt to be sent to the API.
// Model is the ID of the OpenAI model used for the chat completion.
// Messages is an array of Message objects representing the conversation.
type ChatPrompt struct {
	Model    string     `json:"model"`
	Messages []*Message `json:"messages"`
}

// Choice represents a single choice or response from the API.
// Message is a Message object representing the response message.
// FinishReason is the reason the API ended the message (e.g., "stop", "max_tokens", or "temperature").
// Index is the index of the choice in the response.
type Choice struct {
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
	Index        int     `json:"index"`
}

// Usage represents the token usage for a chat completion.
// PromptTokens is the number of tokens in the prompt.
// CompletionTokens is the number of tokens in the completion.
// TotalTokens is the total number of tokens used in the API call.
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// ChatCompletion represents the API response for a chat completion.
// ID is the unique identifier for the completion.
// Object is the type of object returned by the API, usually "text_completion".
// Created is a Unix timestamp of when the completion was created.
// Model is the ID of the OpenAI model used for the completion.
// Usage is a Usage object representing the token usage for the completion.
// Choices is an array of Choice objects representing the response choices.
type ChatCompletion struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Model   string   `json:"model"`
	Usage   Usage    `json:"usage"`
	Choices []Choice `json:"choices"`
}
