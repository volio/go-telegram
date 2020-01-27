package sender

import (
	"errors"

	"github.com/volio/go-telegram/config"
	"github.com/volio/go-telegram/sender/client"
	"github.com/volio/go-telegram/sender/envelop"
)

type Sender interface {
	SendMessage(msg envelop.Message) error
	SendText(msg envelop.TextMessage) error
	SendSticker(msg envelop.StickerMessage) error
}

type sender struct {
	client client.Client
}

func (s *sender) SendMessage(msg envelop.Message) error {
	switch m := msg.(type) {
	case envelop.TextMessage:
		return s.SendText(m)
	case envelop.StickerMessage:
		return s.SendSticker(m)
	case envelop.PhotoMessage:
		return s.SendPhoto(m)
	default:
		return errors.New("unknown message type")
	}
}

func (s *sender) SendText(msg envelop.TextMessage) error {
	return s.client.DoPost("sendMessage", msg.Request())
}

func (s *sender) SendSticker(msg envelop.StickerMessage) error {
	return s.client.DoPost("sendSticker", msg.Request())
}

func (s *sender) SendPhoto(msg envelop.PhotoMessage) error {
	return s.client.DoPost("sendPhoto", msg.Request())
}

func NewSender(cfg *config.BotConfig) Sender {
	c := client.NewClient(cfg)
	return &sender{
		client: c,
	}
}
