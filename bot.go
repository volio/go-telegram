package telegram

import (
	"fmt"
)

type Bot struct {
	client *botClient
}

func (b *Bot) GetMe() (*User, error) {
	var r GetMeReply
	if err := b.client.DoGet("getMe", nil, &r); err != nil {
		return nil, err
	}
	if !r.OK {
		return nil, fmt.Errorf("do req failed, err code: %v, description: %v", r.ErrorCode, r.Description)
	}
	return &r.Result, nil
}

func (b *Bot) SendText(msg *RTextMessage) (*Message, error) {
	var r MessageReply
	if err := b.client.DoPost("sendMessage", msg, &r); err != nil {
		return nil, err
	}
	if !r.OK {
		return nil, fmt.Errorf("do req failed, err code: %v, description: %v", r.ErrorCode, r.Description)
	}
	return &r.Result, nil
}

func (b *Bot) SendSticker(msg *RStickerMessage) (*Message, error) {
	var r MessageReply
	if err := b.client.DoPost("sendSticker", msg, &r); err != nil {
		return nil, err
	}
	if !r.OK {
		return nil, fmt.Errorf("do req failed, err code: %v, description: %v", r.ErrorCode, r.Description)
	}
	return &r.Result, nil
}

func (b *Bot) SendPhoto(msg *RPhotoMessage) (*Message, error) {
	var r MessageReply
	if err := b.client.DoPost("sendPhoto", msg, &r); err != nil {
		return nil, err
	}
	if !r.OK {
		return nil, fmt.Errorf("do req failed, err code: %v, description: %v", r.ErrorCode, r.Description)
	}
	return &r.Result, nil
}

func (b *Bot) ForwardMessage(msg *RForwardMessage) (*Message, error) {
	var r MessageReply
	if err := b.client.DoPost("forwardMessage", msg, &r); err != nil {
		return nil, err
	}
	if !r.OK {
		return nil, fmt.Errorf("do req failed, err code: %v, description: %v", r.ErrorCode, r.Description)
	}
	return &r.Result, nil
}

func (b *Bot) SendAudio(msg *RAudioMessage) (*Message, error) {
	var r MessageReply
	if err := b.client.DoPost("sendAudio", msg, &r); err != nil {
		return nil, err
	}
	if !r.OK {
		return nil, fmt.Errorf("do req failed, err code: %v, description: %v", r.ErrorCode, r.Description)
	}
	return &r.Result, nil
}

func (b *Bot) SendMediaGroup(msg *RMediaGroupMessage) ([]Message, error) {
	var r MultiMessageReply
	if err := b.client.DoPost("sendMediaGroup", msg, &r); err != nil {
		return nil, err
	}
	if !r.OK {
		return nil, fmt.Errorf("do req failed, err code: %v, description: %v", r.ErrorCode, r.Description)
	}
	return r.Result, nil
}

func NewBot(cfg *Config) *Bot {
	c := newHttpClient(cfg.Request.RequestTimeout, &cfg.Proxy)
	return &Bot{
		client: newBotClient(cfg.Bot.Key, c),
	}
}
