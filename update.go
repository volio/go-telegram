package telegram

import "strings"

type Reply struct {
	OK          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}

type UpdateReply struct {
	Reply
	Result []*Update `json:"result"`
}

type Update struct {
	UpdateID          int64        `json:"update_id"`
	Message           *Message     `json:"message,omitempty"`
	EditedMessage     *Message     `json:"edited_message,omitempty"`
	ChannelPost       *Message     `json:"channel_post,omitempty"`
	EditedChannelPost *Message     `json:"edited_channel_post,omitempty"`
	InlineQuery       *InlineQuery `json:"inline_query,omitempty"`
}

type Message struct {
	MessageID             int64                 `json:"message_id"`
	From                  User                  `json:"from"`
	Date                  int64                 `json:"date"`
	Chat                  Chat                  `json:"chat"`
	ForwardFrom           *User                 `json:"forward_from,omitempty"`
	ForwardFromChat       *Chat                 `json:"forward_from_chat,omitempty"`
	ForwardFromMessageID  *int64                `json:"forward_from_message_id,omitempty"`
	ForwardSignature      *string               `json:"forward_signature,omitempty"`
	ForwardSenderName     *string               `json:"forward_sender_name,omitempty"`
	ForwardDate           *int64                `json:"forward_date,omitempty"`
	ReplyToMessage        *Message              `json:"reply_to_message,omitempty"`
	EditDate              *int64                `json:"edit_date,omitempty"`
	MediaGroupID          *string               `json:"media_group_id,omitempty"`
	AuthorSignature       *string               `json:"author_signature,omitempty"`
	Text                  *string               `json:"text,omitempty"`
	Entities              []MessageEntity       `json:"entities,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	Audio                 *Audio                `json:"audio,omitempty"`
	Document              *Document             `json:"document,omitempty"`
	Animation             *Animation            `json:"animation,omitempty"`
	Game                  *Game                 `json:"game,omitempty"`
	Photo                 []PhotoSize           `json:"photo,omitempty"`
	Sticker               *Sticker              `json:"sticker,omitempty"`
	Video                 *Video                `json:"video,omitempty"`
	Voice                 *Voice                `json:"voice,omitempty"`
	VideoNote             *VideoNote            `json:"video_note,omitempty"`
	Caption               *string               `json:"caption,omitempty"`
	Contact               *Contact              `json:"contact,omitempty"`
	Location              *Location             `json:"location,omitempty"`
	Venue                 *Venue                `json:"venue,omitempty"`
	Poll                  *Poll                 `json:"poll,omitempty"`
	NewChatMembers        []User                `json:"new_chat_members,omitempty"`
	LeftChatMember        *User                 `json:"left_chat_member,omitempty"`
	NewChatTitle          *string               `json:"new_chat_title,omitempty"`
	NewChatPhoto          []PhotoSize           `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto       *bool                 `json:"delete_chat_photo,omitempty"`
	GroupChatCreated      *bool                 `json:"group_chat_created,omitempty"`
	SuperGroupChatCreated *bool                 `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated    *bool                 `json:"channel_chat_created,omitempty"`
	MigrateToChatID       *int64                `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatID     *int64                `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage         *Message              `json:"pinned_message,omitempty"`
	Invoice               *Invoice              `json:"invoice,omitempty"`
	SuccessfulPayment     *SuccessfulPayment    `json:"successful_payment,omitempty"`
	ConnectedWebsite      *string               `json:"connected_website,omitempty"`
	PassportData          *PassportData         `json:"passport_data,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (m *Message) IsCommand() bool {
	if m.Entities == nil || len(m.Entities) == 0 {
		return false
	}

	entity := m.Entities[0]
	return entity.Offset == 0 && entity.IsCommand()
}

func (m *Message) Command() string {
	command := m.CommandWithAt()

	if i := strings.Index(command, "@"); i != -1 {
		command = command[:i]
	}

	return command
}

func (m *Message) CommandWithAt() string {
	if !m.IsCommand() {
		return ""
	}

	entity := m.Entities[0]
	return (*m.Text)[1:entity.Length]
}

func (m *Message) CommandArguments() string {
	if !m.IsCommand() {
		return ""
	}

	entity := m.Entities[0]
	if len(*m.Text) == entity.Length {
		return ""
	}

	return (*m.Text)[entity.Length+1:]
}

type User struct {
	ID                      int64   `json:"id"`
	IsBot                   bool    `json:"is_bot"`
	FirstName               string  `json:"first_name"`
	LastName                *string `json:"last_name,omitempty"`
	UserName                *string `json:"username,omitempty"`
	LanguageCode            *string `json:"language_code,omitempty"`
	CanJoinGroups           *bool   `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages *bool   `json:"can_read_all_group_messages,omitempty"`
	SupportsInlineQueries   *bool   `json:"supports_inline_queries,omitempty"`
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

type MessageEntity struct {
	Type     string  `json:"type"`
	Offset   int     `json:"offset"`
	Length   int     `json:"length"`
	Url      *string `json:"url,omitempty"`
	User     *User   `json:"user,omitempty"`
	Language *string `json:"language,omitempty"`
}

func (m *MessageEntity) IsCommand() bool {
	return m.Type == "bot_command"
}

type PhotoSize struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int64  `json:"width"`
	Height       int64  `json:"height"`
	FileSize     *int64 `json:"file_size,omitempty"`
}

type Audio struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Duration     int64      `json:"duration"`
	Performer    *string    `json:"performer,omitempty"`
	Title        *string    `json:"title,omitempty"`
	MimeType     *string    `json:"mime_type,omitempty"`
	FileSize     *int64     `json:"file_size,omitempty"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
}

type Document struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
	FileName     *string    `json:"file_name,omitempty"`
	MimeType     *string    `json:"mime_type,omitempty"`
	FileSize     *int64     `json:"file_size,omitempty"`
}

type Animation struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Width        int64      `json:"width"`
	Height       int64      `json:"height"`
	Duration     int64      `json:"duration"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
	FileName     *string    `json:"file_name,omitempty"`
	MimeType     *string    `json:"mime_type,omitempty"`
	FileSize     *int64     `json:"file_size,omitempty"`
}

type Game struct {
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	Photo        []PhotoSize     `json:"photo"`
	Text         *string         `json:"text,omitempty"`
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
	Animation    *Animation      `json:"animation,omitempty"`
}

type Video struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
	MimeType     *string    `json:"mime_type,omitempty"`
	FileSize     *int64     `json:"file_size,omitempty"`
}

type Voice struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Length       int64      `json:"length"`
	Duration     int64      `json:"duration"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
	FileSize     *int64     `json:"file_size,omitempty"`
}

type VideoNote struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Length       int64      `json:"length"`
	Duration     int64      `json:"duration"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
	FileSize     *int64     `json:"file_size,omitempty"`
}

type Contact struct {
	PhoneNumber string  `json:"phone_number"`
	FirstName   string  `json:"first_name"`
	LastName    *string `json:"last_name,omitempty"`
	UserID      *int64  `json:"user_id,omitempty"`
	VCard       *string `json:"vcard,omitempty"`
}

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type Venue struct {
	Location       Location `json:"location"`
	Title          string   `json:"title"`
	Address        string   `json:"address"`
	FoursquareID   *string  `json:"foursquare_id,omitempty"`
	FoursquareType *string  `json:"foursquare_type,omitempty"`
}

type Poll struct {
}

type Invoice struct {
}

type SuccessfulPayment struct {
}

type PassportData struct {
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text                         string        `json:"text"`
	URL                          *string       `json:"url,omitempty"`
	LoginURL                     *string       `json:"login_url,omitempty"`
	CallbackData                 *string       `json:"callback_data,omitempty"`
	SwitchInlineQuery            *string       `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat *string       `json:"switch_inline_query_current_chat,omitempty"`
	CallbackGame                 *CallbackGame `json:"callback_game,omitempty"`
	Pay                          *bool         `json:"pay,omitempty"`
}

type CallbackGame struct {
}

type InlineQuery struct {
	ID       string    `json:"id"`
	From     User      `json:"from"`
	Location *Location `json:"location,omitempty"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
}
