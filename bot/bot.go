package bot

import (
	"errors"
	"fmt"

	"github.com/volio/go-telegram/bot/client"
	"github.com/volio/go-telegram/bot/envelop"
	"github.com/volio/go-telegram/bot/reply"
	"github.com/volio/go-telegram/config"
	"github.com/volio/go-telegram/model"
)

type Bot interface {
	GetMe() (*model.User, error)
	SendMessage(msg envelop.Message) (*model.Message, error)
	SendText(msg envelop.TextMessage) (*model.Message, error)
	SendSticker(msg envelop.StickerMessage) (*model.Message, error)
	ForwardMessage(msg envelop.ForwardMessage) (*model.Message, error)
	SendAudio(msg envelop.AudioMessage) (*model.Message, error)
}

type bot struct {
	client client.Client
}

func (b *bot) GetMe() (*model.User, error) {
	var r reply.GetMeReply
	if err := b.client.DoGet("getMe", nil, &r); err != nil {
		return nil, err
	}
	if !r.OK {
		return nil, fmt.Errorf("do req failed, err code: %v, description: %v", r.ErrorCode, r.Description)
	}
	return &r.Result, nil
}

func (b *bot) SendMessage(msg envelop.Message) (*model.Message, error) {
	switch m := msg.(type) {
	case envelop.TextMessage:
		return b.SendText(m)
	case envelop.StickerMessage:
		return b.SendSticker(m)
	case envelop.PhotoMessage:
		return b.SendPhoto(m)
	default:
		return nil, errors.New("unknown message type")
	}
}

func (b *bot) SendText(msg envelop.TextMessage) (*model.Message, error) {
	var r reply.MessageReply
	if err := b.client.DoPost("sendMessage", msg.Request(), &r); err != nil {
		return nil, err
	}
	if !r.OK {
		return nil, fmt.Errorf("do req failed, err code: %v, description: %v", r.ErrorCode, r.Description)
	}
	return &r.Result, nil
}

func (b *bot) SendSticker(msg envelop.StickerMessage) (*model.Message, error) {
	var r reply.MessageReply
	if err := b.client.DoPost("sendSticker", msg.Request(), &r); err != nil {
		return nil, err
	}
	if !r.OK {
		return nil, fmt.Errorf("do req failed, err code: %v, description: %v", r.ErrorCode, r.Description)
	}
	return &r.Result, nil
}

func (b *bot) SendPhoto(msg envelop.PhotoMessage) (*model.Message, error) {
	var r reply.MessageReply
	if err := b.client.DoPost("sendPhoto", msg.Request(), &r); err != nil {
		return nil, err
	}
	if !r.OK {
		return nil, fmt.Errorf("do req failed, err code: %v, description: %v", r.ErrorCode, r.Description)
	}
	return &r.Result, nil
}

func (b *bot) ForwardMessage(msg envelop.ForwardMessage) (*model.Message, error) {
	var r reply.MessageReply
	if err := b.client.DoPost("forwardMessage", msg.Request(), &r); err != nil {
		return nil, err
	}
	if !r.OK {
		return nil, fmt.Errorf("do req failed, err code: %v, description: %v", r.ErrorCode, r.Description)
	}
	return &r.Result, nil
}

func (b *bot) SendAudio(msg envelop.AudioMessage) (*model.Message, error) {
	var r reply.MessageReply
	if err := b.client.DoPost("sendAudio", msg.Request(), &r); err != nil {
		return nil, err
	}
	if !r.OK {
		return nil, fmt.Errorf("do req failed, err code: %v, description: %v", r.ErrorCode, r.Description)
	}
	return &r.Result, nil
}

func NewBot(cfg *config.BotConfig) Bot {
	c := client.NewClient(cfg)
	return &bot{
		client: c,
	}
}
