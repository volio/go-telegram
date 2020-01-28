package reply

import "github.com/volio/go-telegram/model"

type GetMeReply struct {
	model.Reply
	Result model.User `json:"result"`
}

type MessageReply struct {
	model.Reply
	Result model.Message `json:"result"`
}
