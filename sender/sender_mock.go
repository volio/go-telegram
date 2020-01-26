package sender

import (
	"github.com/stretchr/testify/mock"
	"github.com/volio/go-telegram/sender/envelop"
)

type MockSender struct {
	mock.Mock
}

func (m *MockSender) SendMessage(msg envelop.TextMessage) error {
	args := m.Called(msg)
	return args.Error(0)
}

func (m *MockSender) SendSticker(msg envelop.StickerMessage) error {
	args := m.Called(msg)
	return args.Error(0)
}
