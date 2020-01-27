package request

type TextMessageReq struct {
	ChatID                int64          `json:"chat_id"`
	Text                  string         `json:"text"`
	ParseMode             *string        `json:"parse_mode,omitempty"`
	DisableWebPagePreview *bool          `json:"disable_web_page_preview,omitempty"`
	DisableNotification   *bool          `json:"disable_notification,omitempty"`
	ReplyToMessageID      *int64         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup           ReplyMarkupReq `json:"reply_markup,omitempty"`
}

type StickerMessageReq struct {
	ChatID              int64  `json:"chat_id"`
	Sticker             string `json:"sticker"`
	DisableNotification *bool  `json:"disable_notification,omitempty"`
	ReplyToMessageID    *int64 `json:"reply_to_message_id,omitempty"`
}

type PhotoMessageReq struct {
	ChatID                int64          `json:"chat_id"`
	Photo                 string         `json:"photo"`
	Caption               *string        `json:"caption,omitempty"`
	ParseMode             *string        `json:"parse_mode,omitempty"`
	DisableWebPagePreview *bool          `json:"disable_web_page_preview,omitempty"`
	DisableNotification   *bool          `json:"disable_notification,omitempty"`
	ReplyToMessageID      *int64         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup           ReplyMarkupReq `json:"reply_markup,omitempty"`
}

type ReplyMarkupReq interface {
	t() ReplyMarkupReq
}

type InlineKeyboardMarkupReq struct {
	InlineKeyboard [][]InlineKeyboardButtonReq `json:"inline_keyboard"`
}

func (i *InlineKeyboardMarkupReq) t() ReplyMarkupReq {
	panic("implement me")
}

type InlineKeyboardButtonReq struct {
	Text string  `json:"text"`
	URL  *string `json:"url,omitempty"`
}
