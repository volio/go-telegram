package sender

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/volio/go-telegram/sender/client"
	"github.com/volio/go-telegram/sender/envelop"
)

func TestSender_SendText(t *testing.T) {
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

		err := sender.SendText(msg)
		assert.Nil(t, err)
		c.AssertCalled(t, "DoPost", "sendMessage", msg.Request())
	})
}

func TestSender_SendSticker(t *testing.T) {
	t.Run("send", func(t *testing.T) {
		c := new(client.MockClient)
		c.On("DoPost", mock.Anything, mock.Anything).Return(nil)

		sender := &sender{
			client: c,
		}

		msg := envelop.StickerMessage{
			ChatID:  1,
			Sticker: "hello",
		}

		err := sender.SendSticker(msg)
		assert.Nil(t, err)
		c.AssertCalled(t, "DoPost", "sendSticker", msg.Request())
	})
}

func TestSender_SendPhoto(t *testing.T) {
	t.Run("send", func(t *testing.T) {
		c := new(client.MockClient)
		c.On("DoPost", mock.Anything, mock.Anything).Return(nil)

		sender := &sender{
			client: c,
		}

		msg := envelop.PhotoMessage{
			ChatID: 1,
			Photo:  "hello",
		}

		err := sender.SendPhoto(msg)
		assert.Nil(t, err)
		c.AssertCalled(t, "DoPost", "sendPhoto", msg.Request())
	})
}
