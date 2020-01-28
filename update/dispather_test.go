package update

import (
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/volio/go-telegram/model"
)

type MockHandler struct {
	mock.Mock
}

func (h *MockHandler) Handle(update *model.Update) error {
	args := h.Called(update)
	return args.Error(0)
}

func TestDispatcher_Run(t *testing.T) {
	t.Run("run", func(t *testing.T) {
		h := new(MockHandler)
		h.On("Handle", mock.Anything).Return(nil)

		dispatcher := new(dispatcher)
		ch := make(chan *model.Update, 100)
		qch := make(chan interface{})
		dispatcher.SetHandler(h.Handle)

		go dispatcher.Run(ch, qch)

		u := model.Update{UpdateID: 1}
		// sleep for go routine start
		ch <- &u
		time.Sleep(time.Millisecond * 100)
		close(qch)

		h.AssertCalled(t, "Handle", &u)
	})
}
