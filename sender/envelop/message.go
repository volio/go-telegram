package envelop

type TextMessage struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

type StickerMessage struct {
	ChatID  int64  `json:"chat_id"`
	Sticker string `json:"sticker"`
}
