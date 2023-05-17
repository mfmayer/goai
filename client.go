package gopenai

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type clientOptions struct {
	model string
}

type clientOption func(co *clientOptions)

func WithModel(model string) clientOption {
	return func(co *clientOptions) {
		co.model = model
	}
}

type Client struct {
	model        string
	openaiApiKey string
}

func NewClient(openaiApKey string, opts ...clientOption) *Client {
	options := clientOptions{
		model: "gpt-3.5-turbo",
	}
	for _, opt := range opts {
		opt(&options)
	}
	return &Client{
		openaiApiKey: openaiApKey,
		model:        options.model,
	}
}

func (c *Client) GetChatCompletion(prompt *ChatPrompt) (completion *ChatCompletion, err error) {
	url := "https://api.openai.com/v1/chat/completions"
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + c.openaiApiKey,
	}
	if prompt.Model == "" {
		prompt.Model = c.model
	}
	jsonData, err := json.Marshal(prompt)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	completion = &ChatCompletion{}
	err = json.Unmarshal(body, completion)
	if err != nil {
		return nil, err
	}
	if completion.Error != nil {
		return completion, errors.New(completion.Error.Message)
	}
	return
}
