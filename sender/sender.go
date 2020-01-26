package sender

import (
	"github.com/volio/go-telegram/config"
	"github.com/volio/go-telegram/sender/client"
	"github.com/volio/go-telegram/sender/envelop"
)

type Sender interface {
	SendMessage(chatID int64, text string) error
	SendSticker(chatID int64, sticker string) error
}

type sender struct {
	client client.Client
}

func (s *sender) SendMessage(chatID int64, text string) error {
	textMessage := envelop.TextMessage{
		ChatID: chatID,
		Text:   text,
	}
	return s.client.DoPost("sendMessage", textMessage)
}

func (s *sender) SendSticker(chatID int64, sticker string) error {
	stickerMessage := envelop.StickerMessage{
		ChatID:  chatID,
		Sticker: sticker,
	}
	return s.client.DoPost("sendSticker", stickerMessage)
}

func NewSender(cfg *config.BotConfig) Sender {
	c := client.NewClient(cfg)
	return &sender{
		client: c,
	}
}
