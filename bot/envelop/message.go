package envelop

import "github.com/volio/go-telegram/bot/request"

type Message interface {
}

type ReplyMarkup interface {
	Request() request.ReplyMarkup
}

type TextMessage struct {
	ChatID                int64       `json:"chat_id"`
	Text                  string      `json:"text"`
	ParseMode             string      `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool        `json:"disable_web_page_preview,omitempty"`
	DisableNotification   bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID      int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup           ReplyMarkup `json:"reply_markup,omitempty"`
}

func (m *TextMessage) Request() *request.TextMessage {
	if m == nil {
		return nil
	}
	r := &request.TextMessage{
		ChatID: m.ChatID,
		Text:   m.Text,
	}
	if m.ParseMode != "" {
		r.ParseMode = &m.ParseMode
	}
	if m.DisableNotification {
		r.DisableNotification = &m.DisableNotification
	}
	if m.ReplyToMessageID != 0 {
		r.ReplyToMessageID = &m.ReplyToMessageID
	}
	if m.ReplyMarkup != nil {
		r.ReplyMarkup = m.ReplyMarkup.Request()
	}
	return r
}

type StickerMessage struct {
	ChatID              int64  `json:"chat_id"`
	Sticker             string `json:"sticker"`
	DisableNotification bool   `json:"disable_notification,omitempty"`
	ReplyToMessageID    int64  `json:"reply_to_message_id,omitempty"`
}

func (m *StickerMessage) Request() *request.StickerMessage {
	r := &request.StickerMessage{
		ChatID:  m.ChatID,
		Sticker: m.Sticker,
	}
	if m.DisableNotification {
		r.DisableNotification = &m.DisableNotification
	}
	if m.ReplyToMessageID != 0 {
		r.ReplyToMessageID = &m.ReplyToMessageID
	}
	return r
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

func (i *InlineKeyboardMarkup) Request() request.ReplyMarkup {
	r := new(request.InlineKeyboardMarkup)
	g := make([][]request.InlineKeyboardButton, 0, len(i.InlineKeyboard))
	for _, v := range i.InlineKeyboard {
		b := make([]request.InlineKeyboardButton, 0, len(v))
		for _, button := range v {
			b = append(b, *button.Request())
		}
		g = append(g, b)
	}
	r.InlineKeyboard = g
	return r
}

type InlineKeyboardButton struct {
	Text string `json:"text"`
	URL  string `json:"url,omitempty"`
}

func (b *InlineKeyboardButton) Request() *request.InlineKeyboardButton {
	r := &request.InlineKeyboardButton{
		Text: b.Text,
	}
	if b.URL != "" {
		r.URL = &b.URL
	}
	return r
}

type PhotoMessage struct {
	ChatID                int64       `json:"chat_id"`
	Photo                 string      `json:"photo"`
	Caption               string      `json:"caption,omitempty"`
	ParseMode             string      `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool        `json:"disable_web_page_preview,omitempty"`
	DisableNotification   bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID      int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup           ReplyMarkup `json:"reply_markup,omitempty"`
}

func (m *PhotoMessage) Request() *request.PhotoMessage {
	if m == nil {
		return nil
	}
	r := &request.PhotoMessage{
		ChatID: m.ChatID,
		Photo:  m.Photo,
	}
	if m.Caption != "" {
		r.Caption = &m.Caption
	}
	if m.ParseMode != "" {
		r.ParseMode = &m.ParseMode
	}
	if m.DisableNotification {
		r.DisableNotification = &m.DisableNotification
	}
	if m.ReplyToMessageID != 0 {
		r.ReplyToMessageID = &m.ReplyToMessageID
	}
	if m.ReplyMarkup != nil {
		r.ReplyMarkup = m.ReplyMarkup.Request()
	}
	return r
}
