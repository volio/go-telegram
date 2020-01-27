package sender

import (
	"github.com/volio/go-telegram/config"
	"github.com/volio/go-telegram/sender/client"
	"github.com/volio/go-telegram/sender/envelop"
)

type Sender interface {
	SendMessage(msg envelop.TextMessage) error
	SendSticker(msg envelop.StickerMessage) error
}

type sender struct {
	client client.Client
}

func (s *sender) SendMessage(msg envelop.TextMessage) error {
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
