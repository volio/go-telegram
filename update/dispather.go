package update

import (
	"github.com/volio/go-common/graceful"
	"github.com/volio/go-common/log"
	"github.com/volio/go-telegram/handler"
	"github.com/volio/go-telegram/model"
	"github.com/volio/go-telegram/sender"
	"go.uber.org/zap"
)

type Dispatcher interface {
	Run(ch chan *model.Update, q chan bool)
}

func NewDispatcher(handler handler.UpdateHandler, sender sender.Sender) Dispatcher {
	return &dispatcher{
		handler: handler,
		sender:  sender,
	}
}

type dispatcher struct {
	handler handler.UpdateHandler
	sender  sender.Sender
}

func (d *dispatcher) Run(ch chan *model.Update, q chan bool) {
	for {
		select {
		case v := <-ch:
			graceful.Go(func() {
				if err := d.handler.Handle(v, d.sender); err != nil {
					log.L().With(zap.Error(err)).Error("handle update failed")
				}
			})
		case <-q:
			log.L().Info("exit handler")
			return
		}
	}
}
