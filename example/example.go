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
		config.BotConfig{
			LongPollTimeout: 10 * time.Minute,
			RequestTimeout:  6 * time.Second,
			Key:             "telegram bot key",
		},
	)
	t.OnUpdate(func(update *model.Update) error {
		fmt.Printf("receive update: %+v", update)

		if update.Message == nil || update.Message.Text == nil {
			return nil
		}

		msg := envelop.TextMessage{
			ChatID: update.Message.Chat.ID,
			Text:   *update.Message.Text,
		}

		return t.Bot().SendMessage(msg)
	})
	t.Start()
}
