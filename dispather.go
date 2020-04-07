package telegram

import (
	"sync"

	"github.com/volio/go-common/graceful"
	"github.com/volio/go-common/log"
	"go.uber.org/zap"
)

type HandleFunc func(update *Update) error

func newDispatcher(ch chan *Update) *dispatcher {
	return &dispatcher{
		ch: ch,
	}
}

type dispatcher struct {
	ch           chan *Update
	handlerChain []HandleFunc
	once         sync.Once
}

func (d *dispatcher) RegisterHandler(fn HandleFunc) {
	d.handlerChain = append(d.handlerChain, fn)
}

func (d *dispatcher) Run() {
	d.once.Do(func() {
		for v := range d.ch {
			for _, handleFunc := range d.handlerChain {
				graceful.Go(func() {
					if err := handleFunc(v); err != nil {
						log.L().With(zap.Error(err)).Error("handle update failed")
					}
				})
			}
		}
	})
}
