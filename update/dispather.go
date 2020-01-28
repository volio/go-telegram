package update

import (
	"github.com/volio/go-common/graceful"
	"github.com/volio/go-common/log"
	"github.com/volio/go-telegram/model"
	"go.uber.org/zap"
)

type HandleFunc func(update *model.Update) error

type Dispatcher interface {
	Run(ch chan *model.Update, q chan interface{})
	SetHandler(fn HandleFunc)
}

func NewDispatcher() Dispatcher {
	return new(dispatcher)
}

type dispatcher struct {
	handleFunc HandleFunc
}

func (d *dispatcher) SetHandler(fn HandleFunc) {
	if d.handleFunc != nil {
		panic("duplicate handler registered")
	}
	d.handleFunc = fn
}

func (d *dispatcher) Run(ch chan *model.Update, q chan interface{}) {
	for {
		select {
		case v := <-ch:
			if d.handleFunc != nil {
				graceful.Go(func() {
					if err := d.handleFunc(v); err != nil {
						log.L().With(zap.Error(err)).Error("handle update failed")
					}
				})
			}
		case <-q:
			log.L().Info("exit handler")
			return
		}
	}
}
