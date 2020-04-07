package telegram

import (
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
)

type MockHandler struct {
	mock.Mock
}

func (h *MockHandler) Handle(update *Update) error {
	args := h.Called(update)
	return args.Error(0)
}

func TestDispatcher_Run(t *testing.T) {
	t.Run("run", func(t *testing.T) {
		h := new(MockHandler)
		h.On("Handle", mock.Anything).Return(nil)

		ch := make(chan *Update, 100)
		dispatcher := &dispatcher{ch: ch}
		dispatcher.RegisterHandler(h.Handle)
		dispatcher.RegisterHandler(h.Handle)

		go dispatcher.Run()

		u := Update{UpdateID: 1}
		// sleep for go routine start
		ch <- &u
		time.Sleep(time.Millisecond * 1)

		h.AssertCalled(t, "Handle", &u)
		h.AssertNumberOfCalls(t, "Handle", 2)
	})
}
