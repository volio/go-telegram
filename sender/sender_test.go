package sender

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/volio/go-telegram/sender/client"
	"github.com/volio/go-telegram/sender/envelop"
)

func TestSender_SendMessage(t *testing.T) {
	t.Run("send", func(t *testing.T) {
		c := new(client.MockClient)
		c.On("DoPost", mock.Anything, mock.Anything).Return(nil)

		sender := &sender{
			client: c,
		}

		msg := envelop.TextMessage{
			ChatID: 1,
			Text:   "hello",
		}

		err := sender.SendMessage(msg)
		assert.Nil(t, err)
		c.AssertCalled(t, "DoPost", "sendMessage", msg)
	})
}

func TestSender_SendSticker(t *testing.T) {
	t.Run("send", func(t *testing.T) {
		c := new(client.MockClient)
		c.On("DoPost", mock.Anything, mock.Anything).Return(nil)

		sender := &sender{
			client: c,
		}

		msg := envelop.TextMessage{
			ChatID: 1,
			Text:   "hello",
		}

		err := sender.SendMessage(msg)
		assert.Nil(t, err)
		c.AssertCalled(t, "DoPost", "sendMessage", msg)
	})
}
