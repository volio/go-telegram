package bot

import (
	"errors"

	"github.com/volio/go-telegram/bot/client"
	"github.com/volio/go-telegram/bot/envelop"
	"github.com/volio/go-telegram/config"
)

type Bot interface {
	SendMessage(msg envelop.Message) error
	SendText(msg envelop.TextMessage) error
	SendSticker(msg envelop.StickerMessage) error
}

type bot struct {
	client client.Client
}

func (b *bot) SendMessage(msg envelop.Message) error {
	switch m := msg.(type) {
	case envelop.TextMessage:
		return b.SendText(m)
	case envelop.StickerMessage:
		return b.SendSticker(m)
	case envelop.PhotoMessage:
		return b.SendPhoto(m)
	default:
		return errors.New("unknown message type")
	}
}

func (b *bot) SendText(msg envelop.TextMessage) error {
	return b.client.DoPost("sendMessage", msg.Request())
}

func (b *bot) SendSticker(msg envelop.StickerMessage) error {
	return b.client.DoPost("sendSticker", msg.Request())
}

func (b *bot) SendPhoto(msg envelop.PhotoMessage) error {
	return b.client.DoPost("sendPhoto", msg.Request())
}

func NewBot(cfg *config.BotConfig) Bot {
	c := client.NewClient(cfg)
	return &bot{
		client: c,
	}
}
