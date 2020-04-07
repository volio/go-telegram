package telegram

type GetMeReply struct {
	Reply
	Result User `json:"result"`
}

type MessageReply struct {
	Reply
	Result Message `json:"result"`
}

type MultiMessageReply struct {
	Reply
	Result []Message `json:"result"`
}
