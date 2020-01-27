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

type ReplyMarkupReq interface {
}

type InlineKeyboardMarkupReq struct {
	InlineKeyboard [][]InlineKeyboardButtonReq `json:"inline_keyboard"`
}

type InlineKeyboardButtonReq struct {
	Text string  `json:"text"`
	URL  *string `json:"url,omitempty"`
}
