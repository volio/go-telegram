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
	"github.com/volio/go-telegram/bot/envelop"
	"github.com/volio/go-telegram/config"
	"github.com/volio/go-telegram/model"
)

func main() {
	t := telegram.NewTelegram(
		config.Config{
			LongPollTimeout: 10 * time.Minute,
			RequestTimeout:  6 * time.Second,
		},
	)
	t.OnUpdate(func(update *model.Update) error {
		fmt.Printf("receive update: %+v\n", update)

		if update.Message == nil || update.Message.Text == nil {
			return nil
		}

		msg := envelop.TextMessage{
			ChatID: update.Message.Chat.ID,
			Text:   *update.Message.Text,
		}

		_, err := t.Bot().SendMessage(msg)
		return err
	})
	t.Start()
}
```