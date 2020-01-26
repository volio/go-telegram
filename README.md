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
	"github.com/volio/go-telegram/config"
	"github.com/volio/go-telegram/model"
	"github.com/volio/go-telegram/sender"
	"github.com/volio/go-telegram/sender/envelop"
)

// implement update handler
type updateHandler struct {
}

func (h *updateHandler) Handle(update *model.Update, sender sender.Sender) error {
	fmt.Printf("receive update: %+v", update)

	if update.Message == nil || update.Message.Text == nil {
		return nil
	}

	msg := envelop.TextMessage{
		ChatID: update.Message.Chat.ID,
		Text:   *update.Message.Text,
	}

	return sender.SendMessage(msg)
}

func main() {
	bot := telegram.NewTelegram(
		config.BotConfig{
			LongPollTimeout: 10 * time.Minute,
			RequestTimeout:  6 * time.Second,
			Key:             "telegram bot key",
		},
		new(updateHandler),
	)
	bot.Start()
}
```