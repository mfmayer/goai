package oai

type Client struct {
	openaiApiKey string
}

func NewClient(openaiApKey string) *Client {
	return &Client{
		openaiApiKey: openaiApKey,
	}
}
