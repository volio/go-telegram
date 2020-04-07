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
