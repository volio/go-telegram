package update

import (
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/volio/go-telegram/handler"
	"github.com/volio/go-telegram/model"
	"github.com/volio/go-telegram/sender"
)

func TestDispatcher_Run(t *testing.T) {
	t.Run("run", func(t *testing.T) {
		s := new(sender.MockSender)
		h := new(handler.MockHandler)
		h.On("Handle", mock.Anything, mock.Anything).Return(nil)

		dispatcher := &dispatcher{
			sender:  s,
			handler: h,
		}
		ch := make(chan *model.Update, 100)
		qch := make(chan interface{})

		go dispatcher.Run(ch, qch)

		u := model.Update{UpdateID: 1}
		// sleep for go routine start
		ch <- &u
		time.Sleep(time.Second)
		close(qch)

		h.AssertCalled(t, "Handle", &u, s)
	})
}
