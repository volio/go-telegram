# Go Telegram Bot Framework

![GitHub](https://img.shields.io/github/license/volio/go-telegram)
[![Travis](https://api.travis-ci.org/volio/go-telegram.svg?branch=master)](https://travis-ci.org/volio/go-telegram)

A simple and powerful golang telegram sdk

## Usage
`go get github.com/volio/go-telegram`

## Example

```go
package main

import (
	"fmt"
	"time"

	"github.com/volio/go-telegram"
)

func main() {
	t := telegram.NewTelegram(
		telegram.Config{
			Request: telegram.RequestConfig{
				LongPollTimeout: 10 * time.Minute,
				RequestTimeout:  6 * time.Second,
			},
			Bot: telegram.BotConfig{
				Key: "",
			},
		},
	)
	t.OnUpdate(func(update *telegram.Update) error {
		fmt.Printf("receive update: %+v\n", update)

		if update.Message == nil || update.Message.Text == nil {
			return nil
		}

		msg := &telegram.RTextMessage{
			ChatID: update.Message.Chat.ID,
			Text:   *update.Message.Text,
		}

		_, err := t.Bot().SendText(msg)
		return err
	})
	t.Start()
}
```