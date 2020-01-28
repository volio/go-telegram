package request

type TextMessage struct {
	ChatID                int64       `json:"chat_id"`
	Text                  string      `json:"text"`
	ParseMode             *string     `json:"parse_mode,omitempty"`
	DisableWebPagePreview *bool       `json:"disable_web_page_preview,omitempty"`
	DisableNotification   *bool       `json:"disable_notification,omitempty"`
	ReplyToMessageID      *int64      `json:"reply_to_message_id,omitempty"`
	ReplyMarkup           ReplyMarkup `json:"reply_markup,omitempty"`
}

type StickerMessage struct {
	ChatID              int64  `json:"chat_id"`
	Sticker             string `json:"sticker"`
	DisableNotification *bool  `json:"disable_notification,omitempty"`
	ReplyToMessageID    *int64 `json:"reply_to_message_id,omitempty"`
}

type PhotoMessage struct {
	ChatID                int64       `json:"chat_id"`
	Photo                 string      `json:"photo"`
	Caption               *string     `json:"caption,omitempty"`
	ParseMode             *string     `json:"parse_mode,omitempty"`
	DisableWebPagePreview *bool       `json:"disable_web_page_preview,omitempty"`
	DisableNotification   *bool       `json:"disable_notification,omitempty"`
	ReplyToMessageID      *int64      `json:"reply_to_message_id,omitempty"`
	ReplyMarkup           ReplyMarkup `json:"reply_markup,omitempty"`
}

type ReplyMarkup interface {
	t() ReplyMarkup
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

func (i *InlineKeyboardMarkup) t() ReplyMarkup {
	panic("implement me")
}

type InlineKeyboardButton struct {
	Text string  `json:"text"`
	URL  *string `json:"url,omitempty"`
}
