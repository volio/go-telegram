package handler

import (
	"github.com/stretchr/testify/mock"
	"github.com/volio/go-telegram/model"
	"github.com/volio/go-telegram/sender"
)

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(update *model.Update, sender sender.Sender) error {
	args := m.Called(update, sender)
	return args.Error(0)
}
