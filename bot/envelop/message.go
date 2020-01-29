package envelop

import (
	"github.com/volio/go-telegram/bot/request"
	"github.com/volio/go-telegram/model"
)

type Message interface {
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
	if m == nil {
		return nil
	}
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

type ReplyMarkup interface {
	Request() *model.InlineKeyboardMarkup
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

func (i *InlineKeyboardMarkup) Request() *model.InlineKeyboardMarkup {
	r := new(model.InlineKeyboardMarkup)
	g := make([][]model.InlineKeyboardButton, 0, len(i.InlineKeyboard))
	for _, v := range i.InlineKeyboard {
		b := make([]model.InlineKeyboardButton, 0, len(v))
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

func (b *InlineKeyboardButton) Request() *model.InlineKeyboardButton {
	if b == nil {
		return nil
	}
	r := &model.InlineKeyboardButton{
		Text: b.Text,
	}
	if b.URL != "" {
		r.URL = &b.URL
	}
	return r
}

type PhotoMessage struct {
	ChatID              int64       `json:"chat_id"`
	Photo               string      `json:"photo"`
	Caption             string      `json:"caption,omitempty"`
	ParseMode           string      `json:"parse_mode,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
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

type ForwardMessage struct {
	ChatID              int64 `json:"chat_id"`
	FromChatID          int64 `json:"from_chat_id"`
	DisableNotification bool  `json:"disable_notification,omitempty"`
	MessageID           int64 `json:"message_id"`
}

func (m *ForwardMessage) Request() *request.ForwardMessage {
	if m == nil {
		return nil
	}
	r := &request.ForwardMessage{
		ChatID:     m.ChatID,
		FromChatID: m.FromChatID,
		MessageID:  m.MessageID,
	}
	if m.DisableNotification {
		r.DisableNotification = &m.DisableNotification
	}
	return r
}

type AudioMessage struct {
	ChatID              int64       `json:"chat_id"`
	Audio               string      `json:"audio"`
	Caption             string      `json:"caption,omitempty"`
	ParseMode           string      `json:"parse_mode,omitempty"`
	Duration            int         `json:"duration,omitempty"`
	Performer           string      `json:"performer,omitempty"`
	Title               string      `json:"title,omitempty"`
	Thumb               string      `json:"thumb,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

func (m *AudioMessage) Request() *request.AudioMessage {
	if m == nil {
		return nil
	}
	r := &request.AudioMessage{
		ChatID: m.ChatID,
		Audio:  m.Audio,
	}
	if m.Caption != "" {
		r.Caption = &m.Caption
	}
	if m.ParseMode != "" {
		r.ParseMode = &m.ParseMode
	}
	if m.Duration != 0 {
		r.Duration = &m.Duration
	}
	if m.Performer != "" {
		r.Performer = &m.Performer
	}
	if m.Title != "" {
		r.Title = &m.Title
	}
	if m.Thumb != "" {
		r.Thumb = &m.Thumb
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

type MediaGroupMessage struct {
	ChatID              int64        `json:"chat_id"`
	Media               []InputMedia `json:"media"`
	DisableNotification bool         `json:"disable_notification,omitempty"`
	ReplyToMessageID    int64        `json:"reply_to_message_id,omitempty"`
}

func (m *MediaGroupMessage) Request() *request.MediaGroupMessage {
	if m == nil {
		return nil
	}
	r := &request.MediaGroupMessage{
		ChatID: m.ChatID,
	}
	medias := make([]request.InputMedia, 0, len(m.Media))
	for _, media := range m.Media {
		medias = append(medias, *media.Request())
	}
	r.Media = medias
	if m.DisableNotification {
		r.DisableNotification = &m.DisableNotification
	}
	if m.ReplyToMessageID != 0 {
		r.ReplyToMessageID = &m.ReplyToMessageID
	}
	return r
}

type InputMedia interface {
	Request() *request.InputMedia
}

type InputMediaPhoto struct {
	Type      string `json:"type"`
	Media     string `json:"media"`
	Caption   string `json:"caption,omitempty"`
	ParseMode string `json:"parse_mode,omitempty"`
}

func (i *InputMediaPhoto) Request() *request.InputMedia {
	if i == nil {
		return nil
	}
	r := &request.InputMedia{
		Type:  i.Type,
		Media: i.Media,
	}
	if i.Caption != "" {
		r.Caption = &i.Caption
	}
	if i.ParseMode != "" {
		r.ParseMode = &i.ParseMode
	}
	return r
}

type InputMediaVideo struct {
	Type              string `json:"type"`
	Media             string `json:"media"`
	Thumb             string `json:"thumb,omitempty"`
	Caption           string `json:"caption,omitempty"`
	ParseMode         string `json:"parse_mode,omitempty"`
	Width             int    `json:"width,omitempty"`
	Height            int    `json:"height,omitempty"`
	Duration          int    `json:"duration,omitempty"`
	SupportsStreaming bool   `json:"supports_streaming,omitempty"`
}

func (i *InputMediaVideo) Request() *request.InputMedia {
	if i == nil {
		return nil
	}
	r := &request.InputMedia{
		Type:  i.Type,
		Media: i.Media,
	}
	if i.Thumb != "" {
		r.Thumb = &i.Thumb
	}
	if i.Caption != "" {
		r.Caption = &i.Caption
	}
	if i.ParseMode != "" {
		r.ParseMode = &i.ParseMode
	}
	if i.Width != 0 {
		r.Width = &i.Width
	}
	if i.Height != 0 {
		r.Height = &i.Height
	}
	if i.Duration != 0 {
		r.Duration = &i.Duration
	}
	if i.SupportsStreaming {
		r.SupportsStreaming = &i.SupportsStreaming
	}
	return r
}
