package gopenai

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type OpenAIClient struct {
	openaiApiKey string
	BaseURL      string
	HTTPClient   *http.Client
}

func NewOpenAIClient(openaiApKey string, baseURL string) *OpenAIClient {
	return &OpenAIClient{
		openaiApiKey: openaiApKey,
		BaseURL:      baseURL,
		HTTPClient:   &http.Client{
			// Timeout: time.Second * 30,
		},
	}
}

func (c *OpenAIClient) sendRequest(method, endpoint string, body interface{}, result interface{}) error {
	// Convert into json
	var buf bytes.Buffer
	if body != nil {
		err := json.NewEncoder(&buf).Encode(body)
		if err != nil {
			return err
		}
	}

	// create the request
	req, err := http.NewRequest(method, c.BaseURL+endpoint, &buf)
	if err != nil {
		return err
	}
	if c.openaiApiKey == "" {
		return errors.New("openaiApiKey is empty")
	}
	req.Header.Add("Authorization", "Bearer "+c.openaiApiKey)
	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	// send the request
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// read the answer
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// if result != nil, try to decode result
	if result != nil {
		err = json.Unmarshal(resBody, result)
		if err != nil {
			return err
		}
	}

	return nil
}
