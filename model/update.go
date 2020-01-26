package model

type UpdateReply struct {
	OK          bool      `json:"ok"`
	Description string    `json:"description"`
	Result      []*Update `json:"result"`
}

type Update struct {
	UpdateID int64    `json:"update_id"`
	Message  *Message `json:"message"`
}

type Message struct {
	MessageID      int64    `json:"message_id"`
	From           User     `json:"from"`
	Chat           Chat     `json:"chat"`
	Date           int64    `json:"date"`
	Text           *string  `json:"text,omitempty"`
	Sticker        *Sticker `json:"sticker,omitempty"`
	NewChatMembers []User   `json:"new_chat_members,omitempty"`
	LeftChatMember *User    `json:"left_chat_member,omitempty"`
	ReplyToMessage *Message `json:"reply_to_message,omitempty"`
	Entities       []Entity `json:"entities,omitempty"`
	PinnedMessage  *Message `json:"pinned_message,omitempty"`
}

type User struct {
	ID           int64   `json:"id"`
	IsBot        bool    `json:"is_bot"`
	FirstName    string  `json:"first_name"`
	LanguageCode *string `json:"language_code,omitempty"`
	LastName     *string `json:"last_name,omitempty"`
	UserName     *string `json:"username,omitempty"`
}

type Chat struct {
	ID        int64   `json:"id"`
	Type      string  `json:"type"`
	Title     *string `json:"title,omitempty"`
	UserName  *string `json:"username,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
}

type Sticker struct {
	Width        int64      `json:"width"`
	Height       int64      `json:"height"`
	IsAnimated   bool       `json:"is_animated"`
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Emoji        *string    `json:"emoji,omitempty"`
	SetName      *string    `json:"set_name,omitempty"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
	FileSize     *int64     `json:"file_size,omitempty"`
}

type Entity struct {
	Type     string  `json:"type"`
	Offset   int64   `json:"offset"`
	Length   int64   `json:"length"`
	Url      *string `json:"url,omitempty"`
	User     *User   `json:"user,omitempty"`
	Language *string `json:"language,omitempty"`
}

type PhotoSize struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int64  `json:"file_size"`
	Width        int64  `json:"width"`
	Height       int64  `json:"height"`
}
