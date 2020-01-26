package handler

import (
	"github.com/volio/go-telegram/model"
	"github.com/volio/go-telegram/sender"
)

type UpdateHandler interface {
	Handle(update *model.Update, sender sender.Sender) error
}
