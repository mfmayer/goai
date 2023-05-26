package gopenai

type Role string

const RoleUser Role = "user"
const RoleAssistant Role = "assistant"
const RoleSystem Role = "system"

// Message represents a message object in a chat conversation.
// Role is the sender's role, either "system", "user", or "assistant".
// Content is the text content of the message.
type Message struct {
	Role    Role   `json:"role"`
	Content string `json:"content"`
}

// ChatPrompt represents a chat prompt to be sent to the API.
// Model is the ID of the OpenAI model used for the chat completion.
// Messages is an array of Message objects representing the conversation.
type ChatPrompt struct {
	// Model name, e.g. "gpt-3.5-turbo"
	Model string `json:"model"`
	// A list of messages describing the conversation so fa
	Messages []*Message `json:"messages"`
	// What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic. OpenAIO generally recommends altering this or top_p but not both.
	Temperature float32 `json:"temperature,omitempty"`
	// An alternative to sampling with temperature, called nucleus sampling, where the model considers the results of the tokens with top_p probability mass. So 0.1 means only the tokens comprising the top 10% probability mass are considered. OpenAI generally recommend altering this or temperature but not both
	TopP float32 `json:"top_p,omitempty"`
	// The maximum number of tokens to generate in the chat completion. The total length of input tokens and generated tokens is limited by the model's context length.
	MaxTokens int `json:"max_tokens,omitempty"`
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far, increasing the model's likelihood to talk about new topics.
	PresencePenalty float32 `json:"presence_penalty,omitempty"`
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far, decreasing the model's likelihood to repeat the same line verbatim.
	FrequencyPenalty float32 `json:"frequency_penalty,omitempty"`
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

type Error struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Code    string `json:"code,omitempty"`
	Param   string `json:"param,omitempty"`
}

// ChatCompletion represents the API response for a chat completion.
// ID is the unique identifier for the completion.
// Object is the type of object returned by the API, usually "text_completion".
// Created is a Unix timestamp of when the completion was created.
// Model is the ID of the OpenAI model used for the completion.
// Usage is a Usage object representing the token usage for the completion.
// Choices is an array of Choice objects representing the response choices.
type ChatCompletion struct {
	Error   *Error   `json:"error,omitempty"`
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Model   string   `json:"model"`
	Usage   Usage    `json:"usage"`
	Choices []Choice `json:"choices"`
}
