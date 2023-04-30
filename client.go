package oai

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client struct {
	openaiApiKey string
}

func NewClient(openaiApKey string) *Client {
	return &Client{
		openaiApiKey: openaiApKey,
	}
}

func (c *Client) GetChatCompletion(prompt *ChatPrompt) (completion *ChatCompletion, err error) {
	url := "https://api.openai.com/v1/chat/completions"
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + c.openaiApiKey,
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
	return
}
