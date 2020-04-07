package telegram

type RTextMessage struct {
	ChatID                int64                  `json:"chat_id"`
	Text                  string                 `json:"text"`
	ParseMode             *string                `json:"parse_mode,omitempty"`
	DisableWebPagePreview *bool                  `json:"disable_web_page_preview,omitempty"`
	DisableNotification   *bool                  `json:"disable_notification,omitempty"`
	ReplyToMessageID      *int64                 `json:"reply_to_message_id,omitempty"`
	ReplyMarkup           *RInlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type RStickerMessage struct {
	ChatID              int64  `json:"chat_id"`
	Sticker             string `json:"sticker"`
	DisableNotification *bool  `json:"disable_notification,omitempty"`
	ReplyToMessageID    *int64 `json:"reply_to_message_id,omitempty"`
}

type RPhotoMessage struct {
	ChatID              int64                  `json:"chat_id"`
	Photo               string                 `json:"photo"`
	Caption             *string                `json:"caption,omitempty"`
	ParseMode           *string                `json:"parse_mode,omitempty"`
	DisableNotification *bool                  `json:"disable_notification,omitempty"`
	ReplyToMessageID    *int64                 `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         *RInlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type RForwardMessage struct {
	ChatID              int64 `json:"chat_id"`
	FromChatID          int64 `json:"from_chat_id"`
	DisableNotification *bool `json:"disable_notification,omitempty"`
	MessageID           int64 `json:"message_id"`
}

type RAudioMessage struct {
	ChatID              int64                  `json:"chat_id"`
	Audio               string                 `json:"audio"`
	Caption             *string                `json:"caption,omitempty"`
	ParseMode           *string                `json:"parse_mode,omitempty"`
	Duration            *int                   `json:"duration,omitempty"`
	Performer           *string                `json:"performer,omitempty"`
	Title               *string                `json:"title,omitempty"`
	Thumb               *string                `json:"thumb,omitempty"`
	DisableNotification *bool                  `json:"disable_notification,omitempty"`
	ReplyToMessageID    *int64                 `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         *RInlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type RMediaGroupMessage struct {
	ChatID              int64         `json:"chat_id"`
	Media               []RInputMedia `json:"media"`
	DisableNotification *bool         `json:"disable_notification,omitempty"`
	ReplyToMessageID    *int64        `json:"reply_to_message_id,omitempty"`
}

type RInputMedia struct {
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

type RInlineKeyboardMarkup struct {
	InlineKeyboard [][]RInlineKeyboardButton `json:"inline_keyboard"`
}

type RInlineKeyboardButton struct {
	Text                         string        `json:"text"`
	URL                          *string       `json:"url,omitempty"`
	LoginURL                     *string       `json:"login_url,omitempty"`
	CallbackData                 *string       `json:"callback_data,omitempty"`
	SwitchInlineQuery            *string       `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat *string       `json:"switch_inline_query_current_chat,omitempty"`
	CallbackGame                 *CallbackGame `json:"callback_game,omitempty"`
	Pay                          *bool         `json:"pay,omitempty"`
}

type RCallbackGame struct {
}
