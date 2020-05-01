package telegram

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDispatcher_Run(t *testing.T) {
	t.Run("run", func(t *testing.T) {
		wg := &sync.WaitGroup{}
		wg.Add(2)

		ch := make(chan *Update, 100)
		u := &Update{UpdateID: 1}

		handler := func(update *Update) error {
			assert.Equal(t, update, u)
			wg.Done()
			return nil
		}

		dispatcher := newDispatcher(ch)
		dispatcher.RegisterHandler(handler)
		dispatcher.RegisterHandler(handler)

		go dispatcher.Run()
		ch <- u

		done := make(chan bool)
		go func() {
			wg.Wait()
			done <- true
		}()

		select {
		case <-done:
		case <-time.After(10 * time.Millisecond):
			assert.Fail(t, "handler not performed")
		}
	})
}
