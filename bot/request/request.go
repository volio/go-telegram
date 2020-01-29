package request

import "github.com/volio/go-telegram/model"

type TextMessage struct {
	ChatID                int64                       `json:"chat_id"`
	Text                  string                      `json:"text"`
	ParseMode             *string                     `json:"parse_mode,omitempty"`
	DisableWebPagePreview *bool                       `json:"disable_web_page_preview,omitempty"`
	DisableNotification   *bool                       `json:"disable_notification,omitempty"`
	ReplyToMessageID      *int64                      `json:"reply_to_message_id,omitempty"`
	ReplyMarkup           *model.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type StickerMessage struct {
	ChatID              int64  `json:"chat_id"`
	Sticker             string `json:"sticker"`
	DisableNotification *bool  `json:"disable_notification,omitempty"`
	ReplyToMessageID    *int64 `json:"reply_to_message_id,omitempty"`
}

type PhotoMessage struct {
	ChatID              int64                       `json:"chat_id"`
	Photo               string                      `json:"photo"`
	Caption             *string                     `json:"caption,omitempty"`
	ParseMode           *string                     `json:"parse_mode,omitempty"`
	DisableNotification *bool                       `json:"disable_notification,omitempty"`
	ReplyToMessageID    *int64                      `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         *model.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type ForwardMessage struct {
	ChatID              int64 `json:"chat_id"`
	FromChatID          int64 `json:"from_chat_id"`
	DisableNotification *bool `json:"disable_notification,omitempty"`
	MessageID           int64 `json:"message_id"`
}

type AudioMessage struct {
	ChatID              int64                       `json:"chat_id"`
	Audio               string                      `json:"audio"`
	Caption             *string                     `json:"caption,omitempty"`
	ParseMode           *string                     `json:"parse_mode,omitempty"`
	Duration            *int                        `json:"duration,omitempty"`
	Performer           *string                     `json:"performer,omitempty"`
	Title               *string                     `json:"title,omitempty"`
	Thumb               *string                     `json:"thumb,omitempty"`
	DisableNotification *bool                       `json:"disable_notification,omitempty"`
	ReplyToMessageID    *int64                      `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         *model.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type MediaGroupMessage struct {
	ChatID              int64        `json:"chat_id"`
	Media               []InputMedia `json:"media"`
	DisableNotification *bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    *int64       `json:"reply_to_message_id,omitempty"`
}

type InputMedia struct {
	Type              string  `json:"type"`
	Media             string  `json:"media"`
	Thumb             *string `json:"thumb,omitempty"`
	Caption           *string `json:"caption,omitempty"`
	ParseMode         *string `json:"parse_mode,omitempty"`
	Width             *int    `json:"width,omitempty"`
	Height            *int    `json:"height,omitempty"`
	Duration          *int    `json:"duration,omitempty"`
	SupportsStreaming *bool   `json:"supports_streaming,omitempty"`
}
