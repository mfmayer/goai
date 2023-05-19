# gopenai

> ℹ️ **NOTE:** This project is currenty in development and incomplete. Please check documentation what parts of the OpenAI API can be used with this module.

A Go module for easy integration of OpenAI API calls in your projects.

[![GoDoc](https://pkg.go.dev/badge/github.com/mfmayer/gopenai)](https://pkg.go.dev/github.com/mfmayer/gopenai)


## Description

`gopenai` is a Go module that encapsulates OpenAI API calls, enabling easy integration of OpenAI services in your projects. This module simplifies sending requests to the OpenAI API and handling responses, allowing you to focus on implementing your artificial intelligence.

At the moment it's far away from covering the whole API. At the moment I've just implemented the parts that I needed for the development of a chatbot.

## Installation

To install `gopenai`, run the following command:

```sh
go get github.com/mfmayer/gopenai
```

## Usage

Import the gopenai package in your project:

```go
import (
    "github.com/mfmayer/gopenai"
)
```

Create a new gopenai.Client with your API key:

```go
client := gopenai.NewClient("your_openai_api_key")
```

**...TO BE CONTINUED...**

## Documentation

For more information on using gopenai, including available features and methods, visit the [official documentation](https://pkg.go.dev/github.com/your_username/gopenai).

## Contributing

Contributions to improve gopenai are welcome. Please submit pull requests or issues through GitHub.

## License

gopenai is released under the MIT License. For more information, see the LICENSE file.

