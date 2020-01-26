package telegram

import (
	"github.com/volio/go-common/graceful"
	"github.com/volio/go-telegram/config"
	"github.com/volio/go-telegram/handler"
	"github.com/volio/go-telegram/model"
	"github.com/volio/go-telegram/poll"
	"github.com/volio/go-telegram/sender"
	"github.com/volio/go-telegram/update"
)

type Telegram interface {
	Start()
	Stop()
}

func NewTelegram(cfg config.BotConfig, handler handler.UpdateHandler) Telegram {
	s := sender.NewSender(&cfg)

	return &telegram{
		poll:       poll.NewPoll(&cfg),
		dispatcher: update.NewDispatcher(handler, s),
		ch:         make(chan *model.Update, 100),
		qch:        make(chan bool),
	}
}

type telegram struct {
	poll       poll.Poll
	dispatcher update.Dispatcher
	ch         chan *model.Update
	qch        chan bool
}

func (t *telegram) Start() {
	graceful.Go(func() {
		t.dispatcher.Run(t.ch, t.qch)
	})
	t.poll.Start(t.ch)
}

func (t *telegram) Stop() {
	t.poll.Stop()
	close(t.qch)
}
